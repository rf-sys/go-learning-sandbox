package middleware

import (
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func LoadMiddlewares(r *chi.Mux) *chi.Mux {
	r.Use(chiMiddleware.RealIP)
	r.Use(loggerMiddleware)
	r.Use(chiMiddleware.Recoverer)

	return r
}
