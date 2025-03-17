package main

import (
	"fmt"
	"net/http"

	"recipes.krogowski.dev/api/internal/models"
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
}
