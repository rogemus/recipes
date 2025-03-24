package main

import (
	"errors"
	"fmt"
	"net/http"

	"recipes.krogowski.dev/internal/models"
	repository "recipes.krogowski.dev/internal/repositories"
	"recipes.krogowski.dev/internal/validator"
)

func (app *application) createRecipeHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Steps       []string `json:"steps"`
		Ingredients []struct {
			IngredientID int64   `json:"ingredient_id"`
			UnitID       int64   `json:"unit_id"`
			Amount       float32 `json:"amount"`
		} `json:"ingredients"`
	}

	if err := app.readJSON(w, r, &input); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	recipe := &models.Recipe{
		Title:       input.Title,
		Steps:       input.Steps,
		Description: input.Description,
	}
	var ingredients []*models.IngredientListItem

	for _, item := range input.Ingredients {
		ingredient := models.IngredientListItem{
			IngredientID: item.IngredientID,
			UnitID:       item.UnitID,
			Amount:       item.Amount,
			RecipeID:     recipe.ID,
		}

		ingredients = append(ingredients, &ingredient)
	}

	models.ValidateRecipe(v, recipe)
	models.ValidateIngredientList(v, ingredients)

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	user := app.contextGetUser(r)
	recipe.UserID = user.ID
	recipe.UserName = user.Name

	if err := app.repos.Recipes.Insert(recipe); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if err := app.repos.IngredientLists.Insert(ingredients, recipe.ID); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	list, err := app.repos.IngredientLists.List(recipe.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/recipes/%d", recipe.ID))

	if err := app.writeJSON(w, http.StatusCreated, envelope{"recipe": recipe, "ingredients": list}, headers); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getRecipeHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.serverErrorResponse(w, r, err)
		return
	}

	recipe, err := app.repos.Recipes.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	ingredients, err := app.repos.IngredientLists.List(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if err := app.writeJSON(w, http.StatusOK, envelope{"recipe": recipe, "ingredients": ingredients}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteRecipeHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.serverErrorResponse(w, r, err)
		return
	}

	user := app.contextGetUser(r)

	if err = app.repos.Recipes.Delete(id, user.ID); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if err = app.writeJSON(w, http.StatusOK, envelope{"msg": "recipe successfully deleted"}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listRecipeHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string
		models.Filters
	}

	v := validator.New()
	qs := r.URL.Query()

	input.Title = app.readString(qs, "title", "")
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)
	input.Filters.Sort = app.readString(qs, "sort", "id")
	input.Filters.SortSafelist = []string{"id", "title", "-id", "-title"}

	if models.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	recipes, metadata, err := app.repos.Recipes.List(input.Title, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if err = app.writeJSON(w, http.StatusOK, envelope{"recipes": recipes, "metadata": metadata}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
