package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
	"recipes.krogowski.dev/api/internal/models"
)

type RecipeRepo struct {
	DB *sql.DB
}

func (r RecipeRepo) Get(recipeID int64) (*models.Recipe, error) {
	if recipeID < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
    SELECT id, created_at, title, description, steps, version FROM recipes
    WHERE id = $1;`
	var recipe models.Recipe

	ctx, cancel := context.WithTimeout(context.Background(), DBRequestTimeout)
	defer cancel()

	err := r.DB.QueryRowContext(ctx, query, recipeID).Scan(
		&recipe.ID,
		&recipe.CreatedAt,
		&recipe.Title,
		&recipe.Description,
		pq.Array(&recipe.Steps),
		&recipe.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &recipe, nil
}

func (r RecipeRepo) Insert(recipe *models.Recipe) error {
	query := `
    INSERT INTO recipes (title, description, steps)
    VALUES ($1, $2, $3)
    RETURNING id, created_at, version;`

	ctx, cancel := context.WithTimeout(context.Background(), DBRequestTimeout)
	defer cancel()

	args := []any{recipe.Title, recipe.Description, pq.Array(recipe.Steps)}

	return r.DB.QueryRowContext(ctx, query, args...).Scan(
		&recipe.ID,
		&recipe.CreatedAt,
		&recipe.Version,
	)
}

func (r RecipeRepo) Update(recipe *models.Recipe) error {
	return nil
}

func (r RecipeRepo) List() ([]models.Recipe, error) {
	return nil, nil
}
