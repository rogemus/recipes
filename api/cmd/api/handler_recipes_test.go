package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"recipes.krogowski.dev/internal/models"
)

// createRecipe
// listRecipes
// deleteRecipe
// getRecipeHandler
func TestGetRecipeHandler(t *testing.T) {
	app, assert := MockApp(t)
	serv := httptest.NewServer(app.routes())
	client := serv.Client()

	res, err := client.Get(serv.URL + "/v1/recipes/1")
	body, err := io.ReadAll(res.Body)
	body = bytes.TrimSpace(body)
	defer res.Body.Close()

	var response struct {
		Data struct {
			Ingredients []models.IngredientListItem `json:"ingredients"`
			Recipe      models.Recipe               `json:"recipe"`
		}
		Error string `json:"error"`
	}
	json.Unmarshal(body, &response)

	assert.Nil(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	assert.Equal(1, response.Data.Recipe.ID)

}
