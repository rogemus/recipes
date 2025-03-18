package repository

import (
	"database/sql"
	"errors"
	"time"
)

type Repos struct {
	IngredientLists IngredientListRepo
	Ingredients     IngredientRepo
	Recipes         RecipeRepo
	Units           UnitRepo
	Users           UserRepo
}

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
	ErrNoDataToInsert = errors.New("no data to insert")
)

const DBRequestTimeout = 3 * time.Second

func New(db *sql.DB) Repos {
	return Repos{
		IngredientLists: IngredientListRepo{DB: db},
		Ingredients:     IngredientRepo{DB: db},
		Recipes:         RecipeRepo{DB: db},
		Units:           UnitRepo{DB: db},
		Users:           UserRepo{DB: db},
	}
}
