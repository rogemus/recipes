package handlers

import (
	"net/http"
	"strconv"

	"recipies.krogowski.dev/internal/core"
	"recipies.krogowski.dev/internal/middleware"
	"recipies.krogowski.dev/internal/repository"
	"recipies.krogowski.dev/internal/tmpl"
)

type recipeHandler struct {
	recipies        repository.RecipeRepository
	ingredientsList repository.IngredientsListRepository
	requestHandler
}

func NewRecipeHandler(
	env core.Env,
	recipies repository.RecipeRepository,
	ingredientsList repository.IngredientsListRepository,
) recipeHandler {
	return recipeHandler{
		recipies:        recipies,
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

	recipe, err := h.recipies.Get(id)

	if err != nil {
		h.serverError(w, r, err)
		return
	}

	ingredientsList, err := h.ingredientsList.List(id)
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	data := tmpl.NewData(r)
	data.Recipe = recipe
	data.IngredientList = ingredientsList

	h.render(w, r, http.StatusOK, "recipe.tmpl", data)
}

func (h *recipeHandler) RegisterRoute(mux *http.ServeMux, midw *middleware.Midw) {
	mux.Handle("GET /recipies/{id}", midw.Dynamic.ThenFunc(h.get))
}
