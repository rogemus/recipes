package main

import (
	"net/http"

	"github.com/justinas/alice"
	"recipies.krogowski.dev/ui"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)
	standard := alice.New(app.logRequest, commonHeader)
	protected := dynamic.Append(app.requireAuth)

	mux.Handle("GET /static/", http.FileServerFS(ui.Files))
	mux.Handle("GET /{$}", dynamic.ThenFunc(app.getHome))
	mux.Handle("GET /recipies/{id}", dynamic.ThenFunc(app.getRecipe))
	mux.Handle("GET /recipies/create", protected.ThenFunc(app.getRecipiesCreate))
	mux.Handle("POST /recipies/create", protected.ThenFunc(app.postRecipesCreate))

	mux.Handle("GET /auth/login", dynamic.ThenFunc(app.getLogin))
	mux.Handle("GET /auth/signup", dynamic.ThenFunc(app.getSignup))
	mux.Handle("POST /auth/login", dynamic.ThenFunc(app.postLogin))
	mux.Handle("POST /auth/signup", dynamic.ThenFunc(app.postSignup))
	mux.Handle("POST /auth/logout", protected.ThenFunc(app.postLogout))

	mux.Handle("GET /usr/profile", protected.ThenFunc(app.getUsrProfile))

	return standard.Then(mux)
}
