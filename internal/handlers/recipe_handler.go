package handlers

import (
	"net/http"
	"strconv"

	"recipes.krogowski.dev/internal/core"
	"recipes.krogowski.dev/internal/middleware"
	"recipes.krogowski.dev/internal/repository"
)

type recipeHandler struct {
	recipes        repository.RecipeRepository
	ingredientsList repository.IngredientsListRepository
	requestHandler
}

func NewRecipeHandler(
	env core.Env,
	recipes repository.RecipeRepository,
	ingredientsList repository.IngredientsListRepository,
) recipeHandler {
	return recipeHandler{
		recipes:        recipes,
		ingredientsList: ingredientsList,
		requestHandler:  requestHandler{Env: env},
	}
}

func (h *recipeHandler) get(w http.ResponseWriter, r *http.Request) {
	paramId := r.PathValue("id")
	id, err := strconv.Atoi(paramId)

	if err != nil {
		// Not found
		h.serverError(w, r, err)
		return
	}

	recipe, err := h.recipes.Get(id)

	if err != nil {
		h.serverError(w, r, err)
		return
	}

	ingredientsList, err := h.ingredientsList.List(id)
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	data := h.Tmpl.NewData(r)
	data.Recipe = recipe
	data.IngredientList = ingredientsList

	h.render(w, r, http.StatusOK, "recipe.tmpl", data)
}

func (h *recipeHandler) RegisterRoute(mux *http.ServeMux, midw *middleware.Midw) {
	mux.Handle("GET /recipes/{id}", midw.Dynamic.ThenFunc(h.get))
}
