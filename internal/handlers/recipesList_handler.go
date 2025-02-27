package handlers

import (
	"net/http"

	"recipies.krogowski.dev/internal/core"
	"recipies.krogowski.dev/internal/middleware"
	"recipies.krogowski.dev/internal/repository"
)

type recipeListHandler struct {
	recipes repository.RecipeRepository
	requestHandler
}

func NewRecipeListHandler(env core.Env, recipeRepo repository.RecipeRepository) recipeListHandler {
	return recipeListHandler{
		recipes:        recipeRepo,
		requestHandler: requestHandler{Env: env},
	}
}

func (h *recipeListHandler) get(w http.ResponseWriter, r *http.Request) {
	recipes, err := h.recipes.List()

	if err != nil {
		h.serverError(w, r, err)
		return
	}

	data := h.Tmpl.NewData(r)
	data.Recipies = recipes
	h.render(w, r, http.StatusOK, "recipesList.tmpl", data)
}

func (h *recipeListHandler) RegisterRoute(mux *http.ServeMux, midw *middleware.Midw) {
	mux.Handle("GET /recipes/list", midw.Dynamic.ThenFunc(h.get))
}
