package handlers

import (
	"net/http"

	"recipes.krogowski.dev/internal/core"
	"recipes.krogowski.dev/internal/middleware"
	"recipes.krogowski.dev/internal/repository"
)

type homeHandler struct {
	recipes repository.RecipeRepository
	requestHandler
}

func NewHomeHandler(env core.Env, recipeRepo repository.RecipeRepository) homeHandler {
	return homeHandler{
		recipes:        recipeRepo,
		requestHandler: requestHandler{Env: env},
	}
}

func (h *homeHandler) get(w http.ResponseWriter, r *http.Request) {
	recipes, err := h.recipes.RandomList(10)

	if err != nil {
		h.serverError(w, r, err)
		return
	}

	data := h.Tmpl.NewData(r)
	data.Recipes = recipes

	h.render(w, r, http.StatusOK, "home.tmpl", data)
}

func (h *homeHandler) RegisterRoute(mux *http.ServeMux, midw *middleware.Midw) {
	mux.Handle("GET /{$}", midw.Dynamic.ThenFunc(h.get))
}
