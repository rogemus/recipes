package repository

import (
	"database/sql"
	"errors"
	"time"
)

type Repos struct {
	Recipes         RecipeRepo
	Units           UnitRepo
	Ingredients     IngredientRepo
	IngredientLists IngredientListRepo
}

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
	ErrNoDataToInsert = errors.New("no data to insert")
)

const DBRequestTimeout = 3 * time.Second

func New(db *sql.DB) Repos {
	return Repos{
		Recipes:         RecipeRepo{DB: db},
		Units:           UnitRepo{DB: db},
		Ingredients:     IngredientRepo{DB: db},
		IngredientLists: IngredientListRepo{DB: db},
	}
}
