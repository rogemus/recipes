package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healtcheckHandler)

	router.HandlerFunc(http.MethodPost, "/v1/recipes", app.createRecipeHandler)
	router.HandlerFunc(http.MethodGet, "/v1/recipes/:id", app.getRecipeHandler)

	return router
}
