package models

import "recipes.krogowski.dev/api/internal/validator"

type IngredientList struct {
	item []IngredientListItem
}

type IngredientListItem struct {
	IngredientID   int64   `json:"ingredient_id"`
	IngredientName string  `json:"ingredient_name"`
	UnitID         int64   `json:"unit_id"`
	UnitName       string  `json:"unit_name"`
	RecipeID       int64   `json:"recipe_id"`
	Amount         float32 `json:"amount"`
}

func ValidateIngredientList(v *validator.Validator, list []*IngredientListItem) {
	v.Check(list != nil, "steps", "must be provided")
	v.Check(len(list) >= 1, "ingredient_list", "must contain at least 1 item")
	v.Check(len(list) <= 20, "ingredient_list", "must not contain more than 20 items")
}
