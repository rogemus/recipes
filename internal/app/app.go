package app

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"recipies.krogowski.dev/internal/core"
	"recipies.krogowski.dev/internal/db"
	"recipies.krogowski.dev/internal/handlers"
	"recipies.krogowski.dev/internal/middleware"
	"recipies.krogowski.dev/internal/repository"
	"recipies.krogowski.dev/internal/tmpl"
	"recipies.krogowski.dev/ui"
)

type app struct {
	srv *http.Server
	env core.Env
}

func New() app {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Flags
	port := flag.String("port", "4848", "server port")
	debug := flag.Bool("debug", false, "debug mode")
	dbPath := flag.String("db", "./recipies.db", "db file with path")
	flag.Parse()

	db, err := db.New(*dbPath)

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	tmplCache, err := tmpl.NewCache()

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	ingredientsRepo := repository.NewIngredientRepository(db)
	ingredientsListRepo := repository.NewIngredientsListRepository(db)
	userRepo := repository.NewUserReposiotry(db)
	unitRepo := repository.NewUnitRepository(db)
	recipeRepo := repository.NewRecipeRepository(db)

	env := core.Env{
		Logger:    logger,
		DebugMode: *debug,
		TmplCache: tmplCache,
	}

	mux := http.NewServeMux()
	midw := middleware.New(env, userRepo)
	midw.Init()

	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	homeHandler := handlers.NewHomeHandler(env, recipeRepo)
	recipeHandler := handlers.NewRecipeHandler(env, recipeRepo, ingredientsListRepo)
	recipeCreateHandler := handlers.NewRecipeCreateHandler(env, recipeRepo, ingredientsListRepo, ingredientsRepo, unitRepo)

	homeHandler.RegisterRoute(mux, midw)
	recipeHandler.RegisterRoute(mux, midw)
	recipeCreateHandler.RegisterRoute(mux, midw)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", *port),
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
		Handler:      midw.StandardChain.Then(mux),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return app{
		srv: srv,
		env: env,
	}
}

func (app *app) Start() {
	app.env.Logger.Info("stating server", "addr", app.srv.Addr, "debugMode", app.env.DebugMode)
	err := app.srv.ListenAndServe()
	app.env.Logger.Error(err.Error())
	os.Exit(1)
}
