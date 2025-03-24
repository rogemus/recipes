package repository

import (
	"context"
	"database/sql"

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

func (m *IngredientRepo) Search(searchQuery string) ([]*models.Ingredient, error) {
	query := `
    SELECT id, name 
    FROM ingredients
    WHERE (to_tsvector('simple', name) @@ plainto_tsquery('simple', $1) OR $1 = '')
    ORDER BY name ASC, id ASC
    LIMIT 5;`

	rows, err := m.DB.Query(query, searchQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ingredients []*models.Ingredient

	for rows.Next() {
		var ingredient models.Ingredient

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
