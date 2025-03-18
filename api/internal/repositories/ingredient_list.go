package repository

import "database/sql"

type IngredientListRepo struct {
	DB *sql.DB
}
