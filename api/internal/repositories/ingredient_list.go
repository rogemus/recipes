package repository

import (
	"context"
	"database/sql"

	"recipes.krogowski.dev/api/internal/models"
)

type IngredientListRepo struct {
	DB *sql.DB
}

func (r *IngredientListRepo) List(recipeID int64) ([]*models.IngredientListItem, error) {
	query := `
    SELECT
      ingredients.id,
      ingredients.name,
      amount,
      units.id,
      units.name,
      recipe_id
    FROM
      ingredients_list
      INNER JOIN units ON ingredients_list.unit_id = units.id
      INNER JOIN ingredients ON ingredients_list.ingredient_id = ingredients.id
    WHERE
      recipe_id = $1
    ORDER BY ingredients.id ASC
    LIMIT 50;`

	ctx, cancel := context.WithTimeout(context.Background(), DBRequestTimeout)
	defer cancel()

	rows, err := r.DB.QueryContext(ctx, query, recipeID)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var ingredientsList []*models.IngredientListItem

	for rows.Next() {
		var item models.IngredientListItem

		err = rows.Scan(
			&item.IngredientID,
			&item.IngredientName,
			&item.Amount,
			&item.UnitID,
			&item.UnitName,
			&item.RecipeID,
		)

		if err != nil {
			return nil, err
		}

		ingredientsList = append(ingredientsList, &item)
	}

	if err = rows.Close(); err != nil {
		return nil, err
	}

	return ingredientsList, nil
}

func (r *IngredientListRepo) Insert(ingredients []*models.IngredientListItem, recipeID int64) error {
	if len(ingredients) == 0 {
		return ErrNoDataToInsert
	}

	ctx, cancel := context.WithTimeout(context.Background(), DBRequestTimeout)
	defer cancel()

	query := `
    INSERT INTO ingredients_list (ingredient_id, unit_id, amount, recipe_id)
    VALUES ($1, $2, $3, $4);`

	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	for _, item := range ingredients {
		args := []any{item.IngredientID, item.UnitID, item.Amount, recipeID}
		_, err = tx.ExecContext(ctx, query, args...)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
