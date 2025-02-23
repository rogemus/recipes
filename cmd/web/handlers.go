package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"recipies.krogowski.dev/internal/models"
	"recipies.krogowski.dev/internal/validator"
)

func (app *application) getHome(w http.ResponseWriter, r *http.Request) {
	recipies, err := app.recipies.RandomList(10)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.Recipies = recipies

	app.render(w, r, http.StatusOK, "home.tmpl", data)
}

func (app *application) getRecipe(w http.ResponseWriter, r *http.Request) {
	paramId := r.PathValue("id")
	id, err := strconv.Atoi(paramId)

	if err != nil {
		app.clientError(w, http.StatusNotFound)
		return
	}

	recipe, err := app.recipies.Get(id)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	ingredientsList, err := app.ingredientsList.List(id)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.Recipe = recipe
	data.IngredientList = ingredientsList

	app.render(w, r, http.StatusOK, "recipe.tmpl", data)
}

type recipieCreateForm struct {
	Title       string
	Description string
	Ingredients []string
	Units       []string
	Amount      []string
	validator.Validator
}

func (app *application) postRecipesCreate(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
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
		data := app.newTemplateData(r)

		ingredients, err := app.ingredients.List()
		if err != nil {
			app.serverError(w, r, err)
			return
		}

		units, err := app.units.List()
		if err != nil {
			app.serverError(w, r, err)
			return
		}

		data.Ingredients = ingredients
		data.Units = units
		data.Form = form

		app.render(w, r, http.StatusUnprocessableEntity, "recipeCreate.tmpl", data)
		return
	}

	userId := app.sessionUserId(r)
	id, err := app.recipies.Insert(form.Title, form.Description, userId)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = app.ingredientsList.Insert(form.Ingredients, form.Units, form.Amount, id)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.SetFlashMsg(r, MsgRecipeCreated)

	pagePath := fmt.Sprintf("/recipies/%d", id)
	http.Redirect(w, r, pagePath, http.StatusSeeOther)
}

func (app *application) getRecipiesCreate(w http.ResponseWriter, r *http.Request) {
	units, err := app.units.List()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	ingredients, err := app.ingredients.List()
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	data := app.newTemplateData(r)
	data.Form = recipieCreateForm{}

	data.Units = units
	data.Ingredients = ingredients

	app.render(w, r, http.StatusOK, "recipeCreate.tmpl", data)
}

type loginForm struct {
	Email    string
	Password string
	validator.Validator
}

func (app *application) getLogin(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = loginForm{}
	app.render(w, r, http.StatusOK, "login.tmpl", data)
}

func (app *application) postLogin(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := loginForm{
		Email:    r.PostForm.Get("email"),
		Password: r.PostForm.Get("password"),
	}

	form.CheckField(validator.NotBlank(form.Email), "email", validator.FieldErr.ErrNotBlank())
	form.CheckField(validator.Matches(form.Email, validator.EmailRX), "email", validator.FieldErr.ErrNotEmail())
	form.CheckField(validator.NotBlank(form.Password), "password", validator.FieldErr.ErrNotBlank())

	if !form.IsValid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, r, http.StatusUnprocessableEntity, "login.tmpl", data)
		return
	}

	id, err := app.users.Authenticate(form.Email, form.Password)

	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			form.AddFormError(validator.FormErros.ErrInvalidCredentials())
			data := app.newTemplateData(r)
			data.Form = form
			app.render(w, r, http.StatusUnprocessableEntity, "login.tmpl", data)
			return
		}

		app.serverError(w, r, err)
		return
	}

	err = app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	user, err := app.users.Get(id)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.sessionManager.Put(r.Context(), userIdSessionKey, id)
	app.sessionManager.Put(r.Context(), userNameSessionKey, user.Name)

	app.SetFlashMsg(r, MsgUserAuthenticeted)
	http.Redirect(w, r, "/recipies/create", http.StatusSeeOther)
}

type signupForm struct {
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	validator.Validator
}

func (app *application) postSignup(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := signupForm{
		Name:            r.PostForm.Get("name"),
		Email:           r.PostForm.Get("email"),
		Password:        r.PostForm.Get("password"),
		PasswordConfirm: r.PostForm.Get("passwordConfirm"),
	}

	form.CheckField(validator.NotBlank(form.Name), "name", validator.FieldErr.ErrNotBlank())
	form.CheckField(validator.NotBlank(form.Email), "email", validator.FieldErr.ErrNotBlank())
	form.CheckField(validator.Matches(form.Email, validator.EmailRX), "email", validator.FieldErr.ErrNotEmail())
	form.CheckField(validator.NotBlank(form.Password), "password", validator.FieldErr.ErrNotBlank())
	form.CheckField(validator.MinChars(form.Password, 8), "password", validator.FieldErr.ErrMinLength(8))
	form.CheckField(validator.NotBlank(form.PasswordConfirm), "passwordConfirm", validator.FieldErr.ErrNotBlank())
	form.CheckField(validator.SameValue(form.Password, form.PasswordConfirm), "passwordConfirm", validator.FieldErr.ErrPassNotSame())

	if !form.IsValid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, r, http.StatusUnprocessableEntity, "signup.tmpl", data)
		return
	}

	err = app.users.Insert(form.Name, form.Email, form.Password)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.SetFlashMsg(r, MsgUserCreated)
	http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
}

func (app *application) getSignup(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = signupForm{}
	app.render(w, r, http.StatusOK, "signup.tmpl", data)
}

func (app *application) postLogout(w http.ResponseWriter, r *http.Request) {
	err := app.sessionManager.RenewToken(r.Context())

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.sessionManager.Remove(r.Context(), "authenticatedUserID")
	app.SetFlashMsg(r, MsgLogout)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) getUsrProfile(w http.ResponseWriter, r *http.Request) {
	userId := app.sessionUserId(r)

	user, err := app.users.Get(userId)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.User = user

	app.render(w, r, http.StatusOK, "usrProfile.tmpl", data)
}
