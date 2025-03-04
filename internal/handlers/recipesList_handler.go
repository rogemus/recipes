package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"recipes.krogowski.dev/internal/core"
	"recipes.krogowski.dev/internal/middleware"
	"recipes.krogowski.dev/internal/repository"
	"recipes.krogowski.dev/internal/validator"
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

const pageSize = 25

func (h *recipeListHandler) get(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	page := r.URL.Query().Get("page")
	order := r.URL.Query().Get("order")

	if page == "" {
		page = "1"
	}

	if order == "" || (order != "asc" && order != "desc") {
		order = "asc"
	}

	pageNumber, err := strconv.Atoi(page)
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	if pageNumber <= 0 {
		h.serverError(w, r, errors.New("pageNumber must be at least 1"))
		return
	}

	pagination, err := h.recipes.Pagination(pageNumber, pageSize)
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	recipes, err := h.recipes.List(query, pageNumber, pageSize, order)
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	data := h.Tmpl.NewData(r)
	data.Recipes = recipes
	data.Pagination = pagination

	h.render(w, r, http.StatusOK, "recipesList.tmpl", data)
}

func (h *recipeListHandler) autocomplete(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data struct {
		Query string `json:"query"`
	}

	if err := decoder.Decode(&data); err != nil {
		// bad request
		h.serverError(w, r, err)
		return
	}

	recipes, err := h.recipes.Search(data.Query)

	if err != nil {
		h.serverError(w, r, err)
		return
	}

	json, err := json.Marshal(recipes)
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func (h *recipeListHandler) search(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		h.serverError(w, r, err)
		return
	}

	form := searchForm{
		Query: r.PostForm.Get("query"),
	}

	form.CheckField(validator.NotBlank(form.Query), "query", validator.FieldErr.ErrNotBlank())
	form.CheckField(validator.MinChars(form.Query, 3), "query", validator.FieldErr.ErrMinLength(3))

	if !form.IsValid() {
		recipes, err := h.recipes.RandomList(10)

		if err != nil {
			h.serverError(w, r, err)
			return
		}

		data := h.Tmpl.NewData(r)
		data.Form = form
		data.Recipes = recipes

		h.render(w, r, http.StatusBadRequest, "home.tmpl", data)
		return
	}

	params := url.Values{}
	params.Add("query", form.Query)
	redirectURL := "/recipes/list?" + params.Encode()

	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

func (h *recipeListHandler) RegisterRoute(mux *http.ServeMux, midw *middleware.Midw) {
	mux.Handle("GET /recipes/list", midw.Dynamic.ThenFunc(h.get))
	mux.Handle("POST /recipes-search", midw.Dynamic.ThenFunc(h.search))
	// TODO add proper middlewares
	mux.HandleFunc("POST /recipes-autocomplete", h.autocomplete)
}
