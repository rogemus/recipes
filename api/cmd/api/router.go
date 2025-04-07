package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/api/v1/healthcheck", app.healtcheckHandler)

	router.HandlerFunc(http.MethodPost, "/api/v1/recipes", app.protected(app.createRecipeHandler))
	router.HandlerFunc(http.MethodGet, "/api/v1/recipes/:id", app.getRecipeHandler)
	router.HandlerFunc(http.MethodGet, "/api/v1/recipes", app.listRecipeHandler)
	router.HandlerFunc(http.MethodDelete, "/api/v1/recipes/:id", app.protected(app.deleteRecipeHandler))

	router.HandlerFunc(http.MethodPost, "/api/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/api/v1/users/activated", app.activateUserHandler)

	router.HandlerFunc(http.MethodPost, "/api/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	return app.recoverPanic(app.logRequest(app.authenticate(router)))
}
