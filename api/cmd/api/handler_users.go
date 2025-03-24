package main

import (
	"errors"
	"net/http"

	"recipes.krogowski.dev/internal/models"
	repository "recipes.krogowski.dev/internal/repositories"
	"recipes.krogowski.dev/internal/validator"
)

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := app.readJSON(w, r, &input); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := &models.User{
		Name:      input.Name,
		Email:     input.Email,
		Activated: false,
	}

	if err := user.Password.Set(input.Password); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	v := validator.New()

	if models.ValidateUser(v, user); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	if err := app.repos.Users.Insert(user); err != nil {
		switch {
		case errors.Is(err, repository.ErrDuplicateEmail):
			v.AddError("email", "a user with this email address already exists")
			app.failedValidationResponse(w, r, v.Errors)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// TODO: send token to user
	token, err := app.repos.Tokens.New(user.ID, app.config.tokens.activationTokenDuration, models.ScopeActivation)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.logger.Info("User activation token", "token", token)

	if err := app.writeJSON(w, http.StatusCreated, envelope{"user": user}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) activateUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		TokenPlaintext string `json:"token"`
	}

	if err := app.readJSON(w, r, &input); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	if models.ValidateTokenPlaintext(v, input.TokenPlaintext); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	user, err := app.repos.Users.GetForToken(models.ScopeActivation, input.TokenPlaintext)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrRecordNotFound):
			v.AddError("token", "invalid or expired activation token")
			app.failedValidationResponse(w, r, v.Errors)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	user.Activated = true

	if err := app.repos.Users.Update(user); err != nil {
		switch {
		case errors.Is(err, repository.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	if err = app.repos.Tokens.DeleteAllForUser(models.ScopeActivation, user.ID); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if err := app.writeJSON(w, http.StatusCreated, envelope{"user": user}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
