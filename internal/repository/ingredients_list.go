package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"recipies.krogowski.dev/internal/models"
)

type ingredientsListRepo struct {
	DB *sql.DB
}

type IngredientsListRepository interface {
	List(recipeId int) ([]models.IngredientsListItem, error)
	Insert(ingredient_ids, unit_ids, amounts []string, recipe_id int) error
}

func NewIngredientsListRepository(db *sql.DB) IngredientsListRepository {
	return &ingredientsListRepo{DB: db}
}

func (m *ingredientsListRepo) List(recipeId int) ([]models.IngredientsListItem, error) {
	stmt := `SELECT
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
    recipe_id = ?;
  `

	rows, err := m.DB.Query(stmt, recipeId)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	ingredientsList := make([]models.IngredientsListItem, 0)

	for rows.Next() {
		var i models.IngredientsListItem

		err = rows.Scan(&i.IngredientID, &i.IngredientName, &i.Amount, &i.UnitID, &i.UnitName, &i.RecipeID)

		if err != nil {
			return nil, err
		}

		ingredientsList = append(ingredientsList, i)
	}

	return ingredientsList, nil
}

func (m *ingredientsListRepo) Insert(ingredient_ids, unit_ids, amounts []string, recipe_id int) error {
	values := make([]string, 0)

	for i, _ := range ingredient_ids {
		unit := unit_ids[i]
		ingredient := ingredient_ids[i]
		amount := amounts[i]
		value := fmt.Sprintf("(%s, %s, %s, %d)", ingredient, unit, amount, recipe_id)
		values = append(values, value)
	}

	valuesStr := strings.Join(values[:], ",")
	stmt := fmt.Sprintf("INSERT INTO ingredients_list (ingredient_id, unit_id, amount, recipe_id) VALUES %s", valuesStr)
	_, err := m.DB.Exec(stmt)

	if err != nil {
		return err
	}

	return nil
}
