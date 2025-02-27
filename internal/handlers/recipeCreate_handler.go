package handlers

import (
	"fmt"
	"net/http"

	"recipies.krogowski.dev/internal/consts"
	"recipies.krogowski.dev/internal/core"
	"recipies.krogowski.dev/internal/middleware"
	"recipies.krogowski.dev/internal/repository"
	"recipies.krogowski.dev/internal/validator"
)

type recipeCreateHandler struct {
	recipes         repository.RecipeRepository
	ingredients     repository.IngredientRepository
	ingredientsList repository.IngredientsListRepository
	units           repository.UnitRepository
	requestHandler
}

func NewRecipeCreateHandler(
	env core.Env,
	recipes repository.RecipeRepository,
	ingredientsList repository.IngredientsListRepository,
	ingredients repository.IngredientRepository,
	units repository.UnitRepository,
) recipeCreateHandler {
	return recipeCreateHandler{
		recipes:         recipes,
		ingredients:     ingredients,
		ingredientsList: ingredientsList,
		units:           units,
		requestHandler:  requestHandler{Env: env},
	}
}

type recipieCreateForm struct {
	Title       string
	Description string
	Ingredients []string
	Units       []string
	Amount      []string
	validator.Validator
}

func (h *recipeCreateHandler) post(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		// bad request
		h.serverError(w, r, err)
		return
	}

	form := recipieCreateForm{
		Title:       r.PostForm.Get("title"),
		Description: r.PostForm.Get("description"),
		Units:       r.PostForm["unit_id"],
		Amount:      r.PostForm["amount"],
		Ingredients: r.PostForm["ingredient_id"],
	}

	form.CheckField(validator.NotBlank(form.Title), "title", validator.FieldErr.ErrNotBlank())
	form.CheckField(validator.NotBlank(form.Description), "description", validator.FieldErr.ErrNotBlank())

	for i, _ := range form.Ingredients {
		unit := form.Units[i]
		ingredient := form.Ingredients[i]
		amount := form.Amount[i]

		index := fmt.Sprintf("i-%d", i)
		form.CheckField(validator.NotBlank(amount), index, validator.FieldErr.ErrNotBlank())
		form.CheckField(validator.NotBlank(ingredient), index, validator.FieldErr.ErrNotBlank())
		form.CheckField(validator.NotBlank(unit), index, validator.FieldErr.ErrNotBlank())
	}

	if !form.IsValid() {
		data := h.Tmpl.NewData(r)

		ingredients, err := h.ingredients.List()
		if err != nil {
			h.serverError(w, r, err)
			return
		}

		units, err := h.units.List()
		if err != nil {
			h.serverError(w, r, err)
			return
		}

		data.Ingredients = ingredients
		data.Units = units
		data.Form = form

		h.render(w, r, http.StatusUnprocessableEntity, "recipeCreate.tmpl", data)
		return
	}

	userId := h.Session.GetUserId(r)
	id, err := h.recipes.Insert(form.Title, form.Description, userId)

	if err != nil {
		h.serverError(w, r, err)
		return
	}

	err = h.ingredientsList.Insert(form.Ingredients, form.Units, form.Amount, id)

	if err != nil {
		h.serverError(w, r, err)
		return
	}

	h.Session.SetFlashMsg(r, consts.MsgRecipeCreated)

	pagePath := fmt.Sprintf("/recipies/%d", id)
	http.Redirect(w, r, pagePath, http.StatusSeeOther)
}

func (h *recipeCreateHandler) get(w http.ResponseWriter, r *http.Request) {
	units, err := h.units.List()
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	ingredients, err := h.ingredients.List()
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	data := h.Tmpl.NewData(r)
	data.Form = recipieCreateForm{}

	data.Units = units
	data.Ingredients = ingredients

	h.render(w, r, http.StatusOK, "recipeCreate.tmpl", data)
}

func (h *recipeCreateHandler) RegisterRoute(mux *http.ServeMux, midw *middleware.Midw) {
	mux.Handle("GET /recipies/create", midw.Protected.ThenFunc(h.get))
	mux.Handle("POST /recipies/create", midw.Protected.ThenFunc(h.post))
}
