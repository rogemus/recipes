package repository

import (
	"database/sql"

	"recipies.krogowski.dev/internal/models"
)

type ingredientRepo struct {
	DB *sql.DB
}

type IngredientRepository interface {
	Insert(name string) error
	Search(query string) ([]models.Ingredient, error)
	List() ([]models.Ingredient, error)
}

func NewIngredientRepository(db *sql.DB) IngredientRepository {
	return &ingredientRepo{DB: db}
}

func (m *ingredientRepo) Insert(name string) error {
	stmt := `INSERT INTO ingredients (name) VALUES (?);`

	_, err := m.DB.Exec(stmt, name)

	if err != nil {
		return err
	}

	return nil
}

func (m *ingredientRepo) Search(query string) ([]models.Ingredient, error) {
	// TODO
	return nil, nil
}

func (m *ingredientRepo) List() ([]models.Ingredient, error) {
	stmt := `SELECT id, name FROM ingredients`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ingredients := make([]models.Ingredient, 0)

	for rows.Next() {
		var i models.Ingredient

		err = rows.Scan(&i.ID, &i.Name)
		if err != nil {
			return nil, err
		}
		ingredients = append(ingredients, i)
	}

	return ingredients, nil
}
