package repository

import (
	"context"
	"database/sql"
	"fmt"

	"recipes.krogowski.dev/internal/models"
)

type IngredientRepo struct {
	DB *sql.DB
}

func (r *IngredientRepo) Insert(ingredientName string) (int64, error) {
	query := `
    INSERT INTO ingredients (name) VALUES ($1)
    RETURNING id;`

	ctx, cancel := context.WithTimeout(context.Background(), DBRequestTimeout)
	defer cancel()

	var ingredientId int64
	err := r.DB.QueryRowContext(ctx, query, ingredientName).Scan(&ingredientId)
	if err != nil {
		return 0, err
	}

	return ingredientId, nil
}

func (r *IngredientRepo) Search(ingredientName string) ([]*models.IngredientSimple, error) {
	query := fmt.Sprintf(`
		SELECT id, name
		FROM ingredients
	  WHERE (to_tsvector('simple', name) @@ to_tsquery('simple', '%s:*') OR $1 = '') 
		ORDER BY name ASC, id ASC
		LIMIT 5;`, ingredientName)

	ctx, cancel := context.WithTimeout(context.Background(), DBRequestTimeout)
	defer cancel()

	args := []any{ingredientName}
	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ingredients []*models.IngredientSimple

	for rows.Next() {
		var ingredient models.IngredientSimple

		if err = rows.Scan(&ingredient.ID, &ingredient.Name); err != nil {
			return nil, err
		}

		ingredients = append(ingredients, &ingredient)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ingredients, nil
}
