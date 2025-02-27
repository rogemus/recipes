package tmpl

import (
	"net/http"
	"time"

	"github.com/justinas/nosurf"
	"recipies.krogowski.dev/internal/models"
)

type TmplData struct {
	CurrentYear     int
	Recipe          models.Recipe
	User            models.User
	Recipies        []models.Recipe
	Ingredients     []models.Ingredient
	Units           []models.Unit
	IngredientList  []models.IngredientsListItem
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
	UserName        string
}

func NewData(r *http.Request) TmplData {
	return TmplData{
		CurrentYear: time.Now().Year(),
		// Flash:           app.sessionManager.PopString(r.Context(), "flash"),
		// IsAuthenticated: app.isAuthenticated(r),
		//   UserName:        app.sessionUserName(r),
		CSRFToken: nosurf.Token(r),
	}
}
