package app

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"recipes.krogowski.dev/internal/core"
	"recipes.krogowski.dev/internal/ctx"
	"recipes.krogowski.dev/internal/db"
	"recipes.krogowski.dev/internal/handlers"
	"recipes.krogowski.dev/internal/middleware"
	"recipes.krogowski.dev/internal/repository"
	"recipes.krogowski.dev/internal/session"
	"recipes.krogowski.dev/internal/tmpl"
	"recipes.krogowski.dev/ui"
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
	dbHost := flag.String("dbHost", "localhost", "db host")
	dbPort := flag.Int("dbPort", 5432, "db host port")
	dbUser := flag.String("dbUser", "recipes", "db user name")
	dbPass := flag.String("dbPass", "Pass1234", "db user password")
	dbName := flag.String("dbName", "recipes", "db name")

	flag.Parse()

	db, err := db.New(
		*dbHost, *dbPort, *dbUser, *dbPass, *dbName,
	)

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	ingredientsRepo := repository.NewIngredientRepository(db)
	ingredientsListRepo := repository.NewIngredientsListRepository(db)
	userRepo := repository.NewUserReposiotry(db)
	unitRepo := repository.NewUnitRepository(db)
	recipeRepo := repository.NewRecipeRepository(db)

	session := session.New(db)
	ctx := ctx.New()
	tmpl := tmpl.New(session, ctx)

	if err = tmpl.NewCache(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// TODO find better way to pass Session
	env := core.Env{
		Session:   session,
		Tmpl:      tmpl,
		Context:   ctx,
		Logger:    logger,
		DebugMode: *debug,
	}

	mux := http.NewServeMux()
	midw := middleware.New(env, userRepo)
	midw.Init()

	thumbnailServer := http.FileServer(http.Dir("./files"))

	mux.Handle("GET /static/", http.FileServerFS(ui.Files))
	mux.Handle("GET /files/", http.StripPrefix("/files", thumbnailServer))

	homeHandler := handlers.NewHomeHandler(env, recipeRepo)
	recipeHandler := handlers.NewRecipeHandler(env, recipeRepo, ingredientsListRepo)
	recipeCreateHandler := handlers.NewRecipeCreateHandler(env, recipeRepo, ingredientsListRepo, ingredientsRepo, unitRepo)
	loginHandler := handlers.NewLoginHandler(env, userRepo)
	logoutHandler := handlers.NewLogoutHandler(env)
	signupHandler := handlers.NewSignupHandler(env, userRepo)
	userProfileHandler := handlers.NewUserProfileHandler(env, userRepo)
	recipesListHandler := handlers.NewRecipeListHandler(env, recipeRepo)

	loginHandler.RegisterRoute(mux, midw)
	logoutHandler.RegisterRoute(mux, midw)
	signupHandler.RegisterRoute(mux, midw)
	userProfileHandler.RegisterRoute(mux, midw)
	homeHandler.RegisterRoute(mux, midw)
	recipeHandler.RegisterRoute(mux, midw)
	recipeCreateHandler.RegisterRoute(mux, midw)
	recipesListHandler.RegisterRoute(mux, midw)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", *port),
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
		Handler:      midw.Standard.Then(mux),
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
