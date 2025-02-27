package repository

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"recipies.krogowski.dev/internal/consts"
	"recipies.krogowski.dev/internal/models"
)

type UserRepository interface {
	Insert(name, email, password string) error
	Authenticate(email, password string) (int, error)
	Exists(id int) (bool, error)
	Get(id int) (models.User, error)
}

type userRepo struct {
	DB *sql.DB
}

func NewUserReposiotry(db *sql.DB) UserRepository {
	return &userRepo{DB: db}
}

func (m *userRepo) Insert(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (name, email, hashed_password) VALUES (?, ?, ?)`
	_, err = m.DB.Exec(stmt, name, email, string(hashedPassword))

	if err != nil {
		// TODO handle dup email
		return err
	}

	return nil
}

func (m *userRepo) Authenticate(email, password string) (int, error) {
	var id int
	var hashedPassword []byte

	stmt := "SELECT id, hashed_password FROM users WHERE email = ?"
	err := m.DB.QueryRow(stmt, email).Scan(&id, &hashedPassword)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, consts.ErrInvalidCredentials
		}

		return 0, err
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, consts.ErrInvalidCredentials
		}

		return 0, err
	}

	return id, nil
}

func (m *userRepo) Exists(id int) (bool, error) {
	var exists bool
	stmt := "SELECT EXISTS(SELECT true FROM users WHERE id = ?)"
	err := m.DB.QueryRow(stmt, id).Scan(&exists)

	return exists, err
}

func (m *userRepo) Get(id int) (models.User, error) {
	var user models.User
	stmt := "SELECT id, name, email FROM users WHERE id = ?"

	err := m.DB.QueryRow(stmt, id).Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
