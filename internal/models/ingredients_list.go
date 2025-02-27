package models

type IngredientsListItem struct {
	IngredientName string
	IngredientID   int
	UnitName       string
	UnitID         int
	RecipeID       int
	Amount         float32
}
