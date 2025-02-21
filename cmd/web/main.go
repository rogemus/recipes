package main

import (
	"flag"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/sqlite3store"
	"github.com/alexedwards/scs/v2"
	"recipies.krogowski.dev/internal/models"
)

type application struct {
	logger          *slog.Logger
	debugMode       bool
	recipies        models.RecipeModelInf
	users           models.UserModelInf
	tmplCache       map[string]*template.Template
	sessionManager  *scs.SessionManager
	ingredientsList models.IngredientsListModelInf
	ingredients     models.IngredientModelInf
	units           models.UnitModelInf
}

func main() {
	// flags
	port := flag.String("port", "4848", "server port")
	debug := flag.Bool("debug", false, "debug mode")
	dbPath := flag.String("db", "./recipies.db", "db file with path")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	db, err := NewDb(*dbPath)

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	tmplCache, err := newTemplateCache()

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	sessionManager := scs.New()
	sessionManager.Store = sqlite3store.New(db)
	sessionManager.Lifetime = 12 * time.Hour

	app := application{
		logger:          logger,
		debugMode:       *debug,
		tmplCache:       tmplCache,
		sessionManager:  sessionManager,
		recipies:        &models.RecipeModel{DB: db},
		users:           &models.UserModel{DB: db},
		units:           &models.UnitModel{DB: db},
		ingredientsList: &models.IngredientsListModel{DB: db},
		ingredients:     &models.IngredientModel{DB: db},
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", *port),
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Info("stating server", "addr", srv.Addr, "debugMode", *debug)
	err = srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
