package repository

import (
	"context"
	"crypto/internal/fips140/sha256"
	"database/sql"
	"errors"
	"time"

	"recipes.krogowski.dev/api/internal/models"
)

type UserRepo struct {
	DB *sql.DB
}

func (r *UserRepo) Insert(user *models.User) error {
	query := `
    INSERT INTO users (name, email, password_hash, activated) 
    VALUES ($1, $2, $3, $4)
    RETURNING id, created_at, version`

	args := []any{user.Name, user.Email, user.Password.Hash, user.Activated}

	ctx, cancel := context.WithTimeout(context.Background(), DBRequestTimeout)
	defer cancel()

	err := r.DB.QueryRowContext(ctx, query, args...).Scan(&user.ID, &user.CreatedAt, &user.Version)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}

	return nil
}

func (r *UserRepo) Update(user *models.User) error {
	query := `
    UPDATE
      users
    SET
      name = $1,
      email = $2,
      password_hash = $3,
      activated = $4,
      version = version + 1
    WHERE
      id = $5
      AND version = $6
    RETURNING
      version;`

	args := []any{
		user.Name,
		user.Email,
		user.Password.Hash,
		user.Activated,
		user.ID,
		user.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), DBRequestTimeout)
	defer cancel()

	err := r.DB.QueryRowContext(ctx, query, args...).Scan(&user.Version)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

func (r *UserRepo) GerByEmail(email string) (*models.User, error) {
	query := `
    SELECT id, created_at, name, email, password_hash, activated, version
    FROM users
    WHERE email = $1`

	var user models.User

	ctx, cancel := context.WithTimeout(context.Background(), DBRequestTimeout)
	defer cancel()

	err := r.DB.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Name,
		&user.Email,
		&user.Password.Hash,
		&user.Activated,
		&user.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil

}

func (r *UserRepo) GetForToken(tokenScope string, tokenPlaintext string) (*models.User, error) {
	query := `
    SELECT
      users.id,
      users.created_at,
      users.name,
      users.email,
      users.password_hash,
      users.activated,
      users.version
    FROM
      users
      INNER JOIN tokens ON users.id = tokens.user_id
    WHERE
      tokens.hash = $1
      AND tokens.scope = $2
      AND tokens.expiry > $3`

	tokenHash := sha256.Sum256([]byte(tokenPlaintext))
	var user models.User

	ctx, cancel := context.WithTimeout(context.Background(), DBRequestTimeout)
	defer cancel()

	args := []any{tokenHash[:], tokenScope, time.Now()}
	err := r.DB.QueryRowContext(ctx, query, args...).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Name,
		&user.Email,
		&user.Password.Hash,
		&user.Activated,
		&user.Version,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}
