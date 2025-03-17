package main

import (
	"errors"
	"fmt"
	"net/http"

	"recipes.krogowski.dev/api/internal/models"
	repository "recipes.krogowski.dev/api/internal/repositories"
	"recipes.krogowski.dev/api/internal/validator"
)

func (app *application) createRecipeHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Steps       []string `json:"steps"`
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

	if models.ValidateRecipe(v, recipe); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	if err := app.repos.Recipes.Insert(recipe); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/recipes/%d", recipe.ID))

	if err := app.writeJSON(w, http.StatusCreated, envelope{"recipe": recipe}, headers); err != nil {
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

	if err := app.writeJSON(w, http.StatusOK, envelope{"recipe": recipe}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteRecipeHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.serverErrorResponse(w, r, err)
		return
	}

	if err = app.repos.Recipes.Delete(id); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if err = app.writeJSON(w, http.StatusOK, envelope{"msg": "recipe successfully deleted"}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
