package main

import (
	"awesomeProject1/db"
	"awesomeProject1/middleware"
	"awesomeProject1/routes"
	"awesomeProject1/runtime"
	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {
	err := runtime.InitRuntimeEnvironment()
	if err != nil {
		log.Fatal().Err(err).Msg("failed setting up an environment")
	}
	// make sure that we do not close DB connection until the program is finished
	defer db.CloseDb(runtime.Database)

	if err := Run(); err != nil {
		log.Fatal().Err(err).Msg("failed running the application")
	}
}

func Run() error {
	log.Info().Msg("Creating router...")
	r := chi.NewRouter()

	log.Info().Msg("Loading middlewares...")
	r = middleware.LoadMiddlewares(r)

	log.Info().Msg("Registering routes...")
	r = routes.LoadRoutes(r)

	log.Info().Msg("Starting up the server...")
	return http.ListenAndServe(":8080", r)
}
