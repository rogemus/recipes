package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"recipies.krogowski.dev/internal/consts"
	"recipies.krogowski.dev/internal/models"
)

type RecipeRepository interface {
	Get(id int) (models.Recipe, error)
	List() ([]models.Recipe, error)
	RandomList(limit int) ([]models.Recipe, error)
	Insert(title, description string, userId int) (int, error)
	Search(query string) ([]models.Recipe, error)
}

type recipeRepo struct {
	DB *sql.DB
}

func NewRecipeRepository(db *sql.DB) RecipeRepository {
	return &recipeRepo{DB: db}
}

func (r *recipeRepo) RandomList(limit int) ([]models.Recipe, error) {
	stmt := `SELECT id, title, description, created FROM recipies ORDER BY RANDOM() LIMIT $1;`

	rows, err := r.DB.Query(stmt, limit)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipies = make([]models.Recipe, 0)

	for rows.Next() {
		var r models.Recipe
		err = rows.Scan(&r.ID, &r.Title, &r.Description, &r.Created)

		if err != nil {
			return nil, err
		}

		recipies = append(recipies, r)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return recipies, nil
}

func (r *recipeRepo) Get(id int) (models.Recipe, error) {
	stmt := `SELECT id, title, description, created FROM recipies WHERE id = $1;`

	recipe := models.Recipe{}

	err := r.DB.QueryRow(stmt, id).Scan(&recipe.ID, &recipe.Title, &recipe.Description, &recipe.Created)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Recipe{}, consts.ErrorNoEntry
		}

		return models.Recipe{}, err
	}

	return recipe, nil
}

func (r *recipeRepo) List() ([]models.Recipe, error) {
	stmt := `SELECT id, title, description, created FROM recipies LIMIT 10`

	rows, err := r.DB.Query(stmt)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipies = make([]models.Recipe, 0)

	for rows.Next() {
		var r models.Recipe
		err = rows.Scan(&r.ID, &r.Title, &r.Description, &r.Created)

		if err != nil {
			return nil, err
		}

		recipies = append(recipies, r)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return recipies, nil
}

func (r *recipeRepo) Insert(title, description string, userId int) (int, error) {
	lastInsertId := 0
	stmt := `INSERT INTO recipies (title, description, user_id) VALUES($1, $2, $3) RETURNING id;`
	err := r.DB.QueryRow(stmt, title, description, userId).Scan(&lastInsertId)

	if err != nil {
		return 0, err
	}

	if lastInsertId == 0 {
		return 0, err
	}

	return int(lastInsertId), nil
}

func (r *recipeRepo) Search(query string) ([]models.Recipe, error) {
	stmt := `SELECT id, title, description, created FROM recipies WHERE LOWER(title) LIKE '%s%s' LIMIT 3;`
	queryStmt := fmt.Sprintf(stmt, query, "%")

	rows, err := r.DB.Query(queryStmt)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipies = make([]models.Recipe, 0)

	for rows.Next() {
		var r models.Recipe
		err = rows.Scan(&r.ID, &r.Title, &r.Description, &r.Created)

		if err != nil {
			return nil, err
		}

		recipies = append(recipies, r)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return recipies, nil
}
