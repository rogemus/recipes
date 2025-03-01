package models

type IngredientsListItem struct {
	IngredientName string  `json:"ingredient_name"`
	IngredientID   int     `json:"ingredient_id"`
	UnitName       string  `json:"unit_name"`
	UnitID         int     `json:"unit_id"`
	RecipeID       int     `json:"recipe_id"`
	Amount         float32 `json:"amount"`
}
