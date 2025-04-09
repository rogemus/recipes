package models

import (
	"time"

	"recipes.krogowski.dev/internal/validator"
)

type Ingredient struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Version   int       `json:"version"`
}

type IngredientSimple struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func ValidateIngredient(v *validator.Validator, ingredient *Ingredient) {
	v.Check(ingredient.Name != "", "name", "must be provided")
	v.Check(len(ingredient.Name) >= 3, "name", "must be at least 3 char long")
	v.Check(len(ingredient.Name) <= 125, "name", "must not be more than 125 bytes long")
}
