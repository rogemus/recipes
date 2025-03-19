package main

import (
	"flag"
	"log/slog"
	"os"
	"time"

	repository "recipes.krogowski.dev/api/internal/repositories"
)

const version = "1.0.0"

type config struct {
	port   int
	env    string
	tokens struct {
		activationTokenDuration time.Duration
	}
	db struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  time.Duration
	}
}

type application struct {
	config config
	logger *slog.Logger
	repos  repository.Repos
}

func main() {
	var cfg config

	/* Flags */
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Enviroment (development|stagging|production)")

	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("RECIPES_DB_DSN"), "PostgreSQL DSN")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.DurationVar(&cfg.db.maxIdleTime, "db-max-idle-time", 15*time.Minute, "PostgreSQL max connection idle time")

	flag.DurationVar(&cfg.tokens.activationTokenDuration, "activation-token-duration", 3*24*time.Hour, "How long token for user activation is valid")

	flag.Parse()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := newDB(cfg)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	repos := repository.New(db)

	app := application{
		config: cfg,
		logger: logger,
		repos:  repos,
	}

	if err = app.serve(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
