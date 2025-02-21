package models

import (
	"database/sql"
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
	Insert(ingredient_ids, unit_ids []int, recipe_id int, amout float32) error
}

func (m *IngredientsListModel) List(recipeId int) ([]IngredientsListItem, error) {
	stmt := `SELECT
    ingredients.id,
    ingredients.name,
    amout,
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

func (m *IngredientsListModel) Insert(ingredient_ids, unit_ids []int, recipe_id int, amout float32) error {
	// TODO add multiple
	ingredient_id := ingredient_ids[0]
	unit_id := unit_ids[0]

	stmt := "INSERT INTO ingredients_list (ingredient_id, unit_id, recipe_id, amout) VALUES (?, ?, ?, ?);"

	_, err := m.DB.Exec(stmt, ingredient_id, unit_id, recipe_id, amout)

	if err != nil {
		return err
	}

	return nil
}
