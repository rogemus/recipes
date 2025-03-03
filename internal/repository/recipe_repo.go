package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"recipes.krogowski.dev/internal/consts"
	"recipes.krogowski.dev/internal/models"
)

type RecipeRepository interface {
	Get(id int) (models.Recipe, error)
	List(query string) ([]models.Recipe, error)
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
	stmt := `SELECT id, title, description, created FROM recipes ORDER BY RANDOM() LIMIT $1;`

	rows, err := r.DB.Query(stmt, limit)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipes = make([]models.Recipe, 0)

	for rows.Next() {
		var r models.Recipe
		err = rows.Scan(&r.ID, &r.Title, &r.Description, &r.Created)

		if err != nil {
			return nil, err
		}

		recipes = append(recipes, r)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return recipes, nil
}

func (r *recipeRepo) Get(id int) (models.Recipe, error) {
	stmt := `SELECT id, title, description, created FROM recipes WHERE id = $1;`

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

func (r *recipeRepo) List(query string) ([]models.Recipe, error) {
	stmt := `SELECT id, title, description, created FROM recipes %s LIMIT 25`

	if query != "" {
		where := fmt.Sprintf("WHERE LOWER(title) LIKE '%s%s'", query, "%")
		stmt = fmt.Sprintf(stmt, where)
	} else {
		stmt = fmt.Sprintf(stmt, "")
	}

	rows, err := r.DB.Query(stmt)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipes = make([]models.Recipe, 0)

	for rows.Next() {
		var r models.Recipe
		err = rows.Scan(&r.ID, &r.Title, &r.Description, &r.Created)

		if err != nil {
			return nil, err
		}

		recipes = append(recipes, r)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return recipes, nil
}

func (r *recipeRepo) Insert(title, description string, userId int) (int, error) {
	lastInsertId := 0
	stmt := `INSERT INTO recipes (title, description, user_id) VALUES($1, $2, $3) RETURNING id;`
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
	stmt := `SELECT id, title, description, created FROM recipes WHERE LOWER(title) LIKE '%s%s' LIMIT 3;`
	queryStmt := fmt.Sprintf(stmt, query, "%")

	rows, err := r.DB.Query(queryStmt)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipes = make([]models.Recipe, 0)

	for rows.Next() {
		var r models.Recipe
		err = rows.Scan(&r.ID, &r.Title, &r.Description, &r.Created)

		if err != nil {
			return nil, err
		}

		recipes = append(recipes, r)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return recipes, nil
}
