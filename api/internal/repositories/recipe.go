package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
	"recipes.krogowski.dev/internal/models"
)

type RecipeRepo struct {
	DB *sql.DB
}

func (r RecipeRepo) Get(recipeID int64) (*models.Recipe, error) {
	if recipeID < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
    SELECT
      recipes.id,
      recipes.created_at,
      recipes.title,
      recipes.description,
      recipes.steps,
      recipes.version,
      users.id,
      users.name
    FROM
      recipes
      LEFT JOIN users ON recipes.id = users.id
    WHERE
      recipes.id = $1;`
	var recipe models.Recipe

	ctx, cancel := context.WithTimeout(context.Background(), DBRequestTimeout)
	defer cancel()

	err := r.DB.QueryRowContext(ctx, query, recipeID).Scan(
		&recipe.ID,
		&recipe.CreatedAt,
		&recipe.Title,
		&recipe.Description,
		pq.Array(&recipe.Steps),
		&recipe.Version,
		&recipe.UserID,
		&recipe.UserName,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &recipe, nil
}

func (r RecipeRepo) Insert(recipe *models.Recipe) error {
	query := `
    INSERT INTO recipes (title, description, steps, user_id)
    VALUES ($1, $2, $3, $4)
    RETURNING id, created_at, version;`

	ctx, cancel := context.WithTimeout(context.Background(), DBRequestTimeout)
	defer cancel()

	args := []any{recipe.Title, recipe.Description, pq.Array(recipe.Steps), recipe.UserID}

	return r.DB.QueryRowContext(ctx, query, args...).Scan(
		&recipe.ID,
		&recipe.CreatedAt,
		&recipe.Version,
	)
}

func (r RecipeRepo) Delete(recipeID int64, userID int64) error {
	if recipeID < 1 {
		return ErrRecordNotFound
	}

	query := `
    DELETE FROM recipes
    WHERE id = $1 AND user_id = $2`

	ctx, cancel := context.WithTimeout(context.Background(), DBRequestTimeout)
	defer cancel()

	result, err := r.DB.ExecContext(ctx, query, recipeID, userID)
	if err != nil {
		return err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (r RecipeRepo) Update(recipe *models.Recipe) error {
	return nil
}

func (r RecipeRepo) List(title string, filters models.Filters) ([]*models.Recipe, models.Metadata, error) {
	query := fmt.Sprintf(`
    SELECT
      count(*) OVER (),
      recipes.id,
      recipes.created_at,
      recipes.title,
      recipes.description,
      recipes.steps,
      recipes.version,
      users.id,
      users.name
    FROM
      recipes
      LEFT JOIN users ON recipes.user_id = users.id
    WHERE (to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 = '')
    ORDER BY recipes.%s %s, recipes.id ASC
    LIMIT $2 OFFSET $3;`, filters.SortColumn(), filters.SortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), DBRequestTimeout)
	defer cancel()

	args := []any{title, filters.Limit(), filters.Offset()}
	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, models.Metadata{}, err
	}
	defer rows.Close()

	totalRecords := 0
	recipes := []*models.Recipe{}

	for rows.Next() {
		var recipe models.Recipe

		err := rows.Scan(
			&totalRecords,
			&recipe.ID,
			&recipe.CreatedAt,
			&recipe.Title,
			&recipe.Description,
			pq.Array(&recipe.Steps),
			&recipe.Version,
			&recipe.UserID,
			&recipe.UserName,
		)
		if err != nil {
			return nil, models.Metadata{}, err
		}

		recipes = append(recipes, &recipe)
	}

	if err = rows.Err(); err != nil {
		return nil, models.Metadata{}, err
	}

	metadata := models.CalculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return recipes, metadata, nil
}
