package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	// Healthcheck
	router.HandlerFunc(http.MethodGet, "/api/v1/healthcheck", app.healtcheckHandler)

	// Recipes
	router.HandlerFunc(http.MethodPost, "/api/v1/recipes", app.protected(app.createRecipeHandler))
	router.HandlerFunc(http.MethodGet, "/api/v1/recipes/:id", app.getRecipeHandler)
	router.HandlerFunc(http.MethodGet, "/api/v1/recipes", app.listRecipeHandler)
	router.HandlerFunc(http.MethodDelete, "/api/v1/recipes/:id", app.protected(app.deleteRecipeHandler))

	// Users
	router.HandlerFunc(http.MethodPost, "/api/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/api/v1/users/activated", app.activateUserHandler)

	// Ingredients
	// TODO: create ingredient

	// Units
	router.HandlerFunc(http.MethodGet, "/api/v1/units", app.listUnitsHandler)

	// Search
	router.HandlerFunc(http.MethodGet, "/api/v1/search/recipes", app.searchRecipeHandler)
	router.HandlerFunc(http.MethodGet, "/api/v1/search/ingredients", app.searchIngredientHandler)

	// Token
	router.HandlerFunc(http.MethodPost, "/api/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	return app.recoverPanic(app.logRequest(app.authenticate(router)))
}
