package repository

import (
	"database/sql"

	"recipes.krogowski.dev/api/internal/models"
)

type RecipeRepo struct {
	DB *sql.DB
}

func (r RecipeRepo) Get(recipeID int64) (*models.Recipe, error) {
	return nil, nil
}

func (r RecipeRepo) Insert(recipe *models.Recipe) error {
	return nil
}

func (r RecipeRepo) Update(recipe *models.Recipe) error {
	return nil
}

func (r RecipeRepo) List() ([]models.Recipe, error) {
	return nil, nil
}
