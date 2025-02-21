package models

import (
	"database/sql"
	"errors"
	"time"
)

type Recipe struct {
	ID          int
	Title       string
	Created     time.Time
	Description string
	UserId      int
}

type RecipeModelInf interface {
	Get(id int) (Recipe, error)
	List() ([]Recipe, error)
	RandomList(limit int) ([]Recipe, error)
	Insert(title, description string, userId int) (int, error)
}

type RecipeModel struct {
	DB *sql.DB
}

func (r *RecipeModel) RandomList(limit int) ([]Recipe, error) {
	stmt := `SELECT id, title, description, created FROM recipies ORDER BY RANDOM() LIMIT ?`

	rows, err := r.DB.Query(stmt, limit)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipies = make([]Recipe, 0)

	for rows.Next() {
		var r Recipe
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

func (r *RecipeModel) Get(id int) (Recipe, error) {
	stmt := `SELECT id, title, description, created FROM recipies WHERE id = ?`

	recipe := Recipe{}

	err := r.DB.QueryRow(stmt, id).Scan(&recipe.ID, &recipe.Title, &recipe.Description, &recipe.Created)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Recipe{}, ErrorNoEntry
		}

		return Recipe{}, err
	}

	return recipe, nil
}

func (r *RecipeModel) List() ([]Recipe, error) {
	stmt := `SELECT id, title, description, instructions, created FROM recipies LIMIT 10`

	rows, err := r.DB.Query(stmt)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipies = make([]Recipe, 0)

	for rows.Next() {
		var r Recipe
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

func (r *RecipeModel) Insert(title, description string, userId int) (int, error) {
	stmt := `INSERT INTO recipies (title, description, user_id) VALUES(?, ?, ?)`
	result, err := r.DB.Exec(stmt, title, description, userId)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}
