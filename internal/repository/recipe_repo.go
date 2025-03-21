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
	List(query string, pageNumber, pageSize int, order string) ([]models.Recipe, error)
	RandomList(limit int) ([]models.Recipe, error)
	Insert(title, description string, userId int, filename, filepath string) (int, error)
	Search(query string) ([]models.Recipe, error)
	Pagination(pageNumber, pageSize int) (models.Pagination, error)
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
	stmt := `SELECT id, title, description, created, thumbnail_name, thumbnail_path FROM recipes WHERE id = $1;`

	recipe := models.Recipe{}

	err := r.DB.QueryRow(stmt, id).Scan(&recipe.ID, &recipe.Title, &recipe.Description, &recipe.Created, &recipe.ThumbnailName, &recipe.ThumbnailPath)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Recipe{}, consts.ErrorNoEntry
		}

		return models.Recipe{}, err
	}

	return recipe, nil
}

func (r *recipeRepo) Pagination(pageNumber, pageSize int) (models.Pagination, error) {
	stmt := `SELECT COUNT (*) FROM recipes;`
	pagination := models.Pagination{
		PageSize: pageSize,
		Page:     pageNumber,
	}

	err := r.DB.QueryRow(stmt).Scan(&pagination.ItemsCount)

	if err != nil {
		return models.Pagination{}, err
	}

	pagination.CountPages(pagination.PageSize)

	return pagination, nil
}

func (r *recipeRepo) List(query string, pageNumber, pageSize int, order string) ([]models.Recipe, error) {
	offset := 0
	stmt := `SELECT id, title, description, created FROM recipes `

	if query != "" {
		stmt = stmt + fmt.Sprintf("WHERE LOWER(title) LIKE '%s%s' ", query, "%")
	}

	if pageNumber > 1 {
		offset = pageSize * pageNumber
	}

	stmt = stmt + fmt.Sprintf(`ORDER BY "created" %s `, order)
	stmt = stmt + fmt.Sprintf(`LIMIT %d OFFSET %d ROWS `, pageSize, offset)
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

func (r *recipeRepo) Insert(title, description string, userId int, filename, filepath string) (int, error) {
	lastInsertId := 0
	stmt := `INSERT INTO recipes (title, description, user_id, thumbnail_name, thumbnail_path) VALUES($1, $2, $3, $4, $5) RETURNING id;`
	err := r.DB.QueryRow(stmt, title, description, userId, filename, filepath).Scan(&lastInsertId)

	if err != nil {
		return 0, err
	}

	if lastInsertId == 0 {
		return 0, err
	}

	return int(lastInsertId), nil
}

func (r *recipeRepo) Search(query string) ([]models.Recipe, error) {
	stmt := `SELECT id, title, description, created FROM recipes WHERE LOWER(title) LIKE '%s%s' LIMIT 5;`
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
