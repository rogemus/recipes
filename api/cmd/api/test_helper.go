package main

import (
	"io"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func MockApp(t *testing.T) (*application, *assert.Assertions) {
	assert := assert.New(t)
	app := &application{
		logger: slog.New(slog.NewTextHandler(io.Discard, nil)),
	}

	return app, assert
}
