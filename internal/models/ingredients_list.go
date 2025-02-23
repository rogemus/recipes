package models

import (
	"database/sql"
	"fmt"
	"strings"
)

type IngredientsListItem struct {
	IngredientName string
	IngredientID   int
	UnitName       string
	UnitID         int
	RecipeID       int
	Amount         float32
}

type IngredientsListModel struct {
	DB *sql.DB
}

type IngredientsListModelInf interface {
	List(recipeId int) ([]IngredientsListItem, error)
	Insert(ingredient_ids, unit_ids, amounts []string, recipe_id int) error
}

func (m *IngredientsListModel) List(recipeId int) ([]IngredientsListItem, error) {
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

	ingredientsList := make([]IngredientsListItem, 0)

	for rows.Next() {
		var i IngredientsListItem

		err = rows.Scan(&i.IngredientID, &i.IngredientName, &i.Amount, &i.UnitID, &i.UnitName, &i.RecipeID)

		if err != nil {
			return nil, err
		}

		ingredientsList = append(ingredientsList, i)
	}

	return ingredientsList, nil
}

func (m *IngredientsListModel) Insert(ingredient_ids, unit_ids, amounts []string, recipe_id int) error {
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
