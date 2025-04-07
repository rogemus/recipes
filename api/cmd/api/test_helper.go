package main

import (
	"io"
	"log/slog"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	repository "recipes.krogowski.dev/internal/repositories"
)

func MockApp(t *testing.T) (*application, *assert.Assertions) {
	assert := assert.New(t)
	db, _, err := sqlmock.New()
	assert.Nil(err)

	repos := repository.New(db)

	app := &application{
		logger: slog.New(slog.NewTextHandler(io.Discard, nil)),
		repos:  repos,
	}

	return app, assert
}
