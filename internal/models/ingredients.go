package models

import (
	"database/sql"
	"time"
)

type Ingredient struct {
	ID      int
	Name    string
	Created time.Time
}

type IngredientModel struct {
	DB *sql.DB
}

type IngredientModelInf interface {
	Insert(name string) error
	Search(query string) ([]Ingredient, error)
	List() ([]Ingredient, error)
}

func (m *IngredientModel) Insert(name string) error {
	stmt := `INSERT INTO ingredients (name) VALUES (?);`

	_, err := m.DB.Exec(stmt, name)

	if err != nil {
		return err
	}

	return nil
}

func (m *IngredientModel) Search(query string) ([]Ingredient, error) {
	// TODO
	return nil, nil
}

func (m *IngredientModel) List() ([]Ingredient, error) {
	stmt := `SELECT id, name FROM ingredients`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ingredients := make([]Ingredient, 0)

	for rows.Next() {
		var i Ingredient

		err = rows.Scan(&i.ID, &i.Name)
		if err != nil {
			return nil, err
		}
		ingredients = append(ingredients, i)
	}

	return ingredients, nil
}
