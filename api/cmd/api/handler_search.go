package main

import (
	"net/http"

	"recipes.krogowski.dev/internal/models"
	"recipes.krogowski.dev/internal/validator"
)

func (app *application) searchRecipeHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string
	}

	qs := r.URL.Query()

	input.Title = app.readString(qs, "title", "")
	recipe := models.RecipeSimple{
		Title: input.Title,
	}

	v := validator.New()
	if models.ValidateRecipeSimple(v, &recipe); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	recipes, err := app.repos.Recipes.ListSimple(input.Title)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	response := envelope{"data": map[string]any{
		"recipes": recipes,
	}}

	if err = app.writeJSON(w, http.StatusOK, response, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
