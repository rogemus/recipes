package handlers

import (
	"net/http"

	"recipies.krogowski.dev/internal/core"
	"recipies.krogowski.dev/internal/middleware"
	"recipies.krogowski.dev/internal/repository"
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
	recipies, err := h.recipes.RandomList(10)

	if err != nil {
		h.serverError(w, r, err)
		return
	}

	data := h.Tmpl.NewData(r)
	data.Recipies = recipies

	h.render(w, r, http.StatusOK, "home.tmpl", data)
}

func (h *homeHandler) RegisterRoute(mux *http.ServeMux, midw *middleware.Midw) {
	mux.Handle("GET /{$}", midw.Dynamic.ThenFunc(h.get))
}
