package models

import (
	"time"

	"recipes.krogowski.dev/api/internal/validator"
)

type Recipe struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	CreatedAt   time.Time `json:"created"`
	Description string    `json:"description"`
	Steps       []string  `json:"steps"`
	Version     int       `json:"version"`
	UserID      int64     `json:"user_id"`
	UserName    string    `json:"user_name"`
}

func ValidateRecipe(v *validator.Validator, recipe *Recipe) {
	v.Check(recipe.Title != "", "title", "must be provided")
	v.Check(len(recipe.Title) <= 125, "title", "must not be more than 125 bytes long")

	v.Check(recipe.Description != "", "description", "must be provided")
	v.Check(len(recipe.Description) <= 500, "description", "must not be more than 500 bytes long")

	v.Check(recipe.Steps != nil, "steps", "must be provided")
	v.Check(len(recipe.Steps) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(recipe.Steps) <= 20, "genres", "must not contain more than 20 genres")
}
