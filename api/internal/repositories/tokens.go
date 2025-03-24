package repository

import (
	"context"
	"database/sql"
	"time"

	"recipes.krogowski.dev/internal/models"
)

type TokenRepo struct {
	DB *sql.DB
}

func (r TokenRepo) New(userID int64, ttl time.Duration, scope string) (*models.Token, error) {
	token := models.GenerateToken(userID, ttl, scope)

	err := r.Insert(token)
	return token, err
}

func (r TokenRepo) Insert(token *models.Token) error {
	query := `
    INSERT INTO tokens (hash, user_id, expiry, scope) 
    VALUES ($1, $2, $3, $4);`

	args := []any{token.Hash, token.UserID, token.Expiry, token.Scope}

	ctx, cancel := context.WithTimeout(context.Background(), DBRequestTimeout)
	defer cancel()

	_, err := r.DB.ExecContext(ctx, query, args...)
	return err
}

func (r TokenRepo) DeleteAllForUser(scope string, userID int64) error {
	query := `
    DELETE FROM tokens 
    WHERE scope = $1 AND user_id = $2`

	ctx, cancel := context.WithTimeout(context.Background(), DBRequestTimeout)
	defer cancel()

	_, err := r.DB.ExecContext(ctx, query, scope, userID)
	return err
}
