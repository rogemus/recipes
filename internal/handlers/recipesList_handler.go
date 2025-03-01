package handlers

import (
	"encoding/json"
	"net/http"

	"recipies.krogowski.dev/internal/core"
	"recipies.krogowski.dev/internal/middleware"
	"recipies.krogowski.dev/internal/repository"
	"recipies.krogowski.dev/internal/validator"
)

type recipeListHandler struct {
	recipes repository.RecipeRepository
	requestHandler
}

type searchForm struct {
	Query string
	validator.Validator
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

func (h *recipeListHandler) search(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data struct {
		Query string `json:"query"`
	}

	if err := decoder.Decode(&data); err != nil {
		// bad request
		h.serverError(w, r, err)
		return
	}

	recipies, err := h.recipes.Search(data.Query)

	if err != nil {
		h.serverError(w, r, err)
		return
	}

	json, err := json.Marshal(recipies)
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func (h *recipeListHandler) RegisterRoute(mux *http.ServeMux, midw *middleware.Midw) {
	mux.Handle("GET /recipies/{$}", midw.Dynamic.ThenFunc(h.get))
	// TODO add proper middlewares
	mux.HandleFunc("POST /recipies-search", h.search)
}
