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
