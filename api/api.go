package api

import (
	"awesomeProject1/handlers"
	"awesomeProject1/middleware"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type Api struct {
	// config api
	getConfigsHandler handlers.GetConfigHandler

	// user api
	getUsersHandler  handlers.GetUsersHandler
	postUsersHandler handlers.PostUsersHandler
}

func (api Api) Startup() error {
	r := chi.NewRouter()

	// middlewares
	r.Use(chiMiddleware.RealIP)
	r.Use(middleware.LoggerMiddleware)
	r.Use(chiMiddleware.Recoverer)

	// routes
	r.Get("/configs", api.getConfigsHandler.Handler())
	r.Get("/users", api.getUsersHandler.Handler())
	r.Post("/users", api.postUsersHandler.Handler())

	return http.ListenAndServe(":8080", r)
}

func NewApi(
	getConfigsHandler handlers.GetConfigHandler,
	getUsersHandler handlers.GetUsersHandler,
	postUsersHandler handlers.PostUsersHandler,
) Api {
	return Api{
		getConfigsHandler: getConfigsHandler,
		getUsersHandler:   getUsersHandler,
		postUsersHandler:  postUsersHandler,
	}
}
