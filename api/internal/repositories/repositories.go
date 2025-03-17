package repository

import (
	"database/sql"
	"errors"
	"time"
)

type Repos struct {
	Recipes RecipeRepo
}

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

const DBRequestTimeout = 3 * time.Second

func New(db *sql.DB) Repos {
	return Repos{
		Recipes: RecipeRepo{DB: db},
	}
}
