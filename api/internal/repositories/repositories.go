package repository

import "database/sql"

type Repos struct {
	Recipes RecipeRepo
}

func New(db *sql.DB) Repos {
	return Repos{
		Recipes: RecipeRepo{DB: db},
	}
}
