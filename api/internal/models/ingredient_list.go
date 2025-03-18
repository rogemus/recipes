package models

type IngredientList struct {
	IngredientID   int64   `json:"ingredient_id"`
	IngredientName string  `json:"ingredient_name"`
	UnitId         int64   `json:"unit_id"`
	UnitName       string  `json:"unit_name"`
	RecipeID       int64   `json:"recipe_id"`
	Amount         float32 `json:"amount"`
}
