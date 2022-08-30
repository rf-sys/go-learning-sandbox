package main

import (
	"awesomeProject1/db"
	"awesomeProject1/log"
	"awesomeProject1/middleware"
	"awesomeProject1/routes"
	"awesomeProject1/runtime"
	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

func main() {
	err := runtime.InitRuntimeEnvironment()
	if err != nil {
		runtime.Logger.Fatal(err, "failed setting up an environment")
	}
	// make sure that we do not close DB connection until the program is finished
	defer db.CloseDb(runtime.Database, runtime.Logger)

	if err := Run(runtime.Logger); err != nil {
		runtime.Logger.Fatal(err, "failed running the application")
	}
}

func Run(logger log.Logger) error {
	logger.Info("Creating router...")
	r := chi.NewRouter()

	logger.Info("Loading middlewares...")
	r = middleware.LoadMiddlewares(r)

	logger.Info("Registering routes...")
	r = routes.LoadRoutes(r)

	logger.Info("Starting up the server...")
	return http.ListenAndServe(":8080", r)
}
