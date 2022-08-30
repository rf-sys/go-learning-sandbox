package routes

import (
	"awesomeProject1/runtime"
	"github.com/go-chi/chi/v5"
)

func LoadRoutes(r *chi.Mux) *chi.Mux {
	r.Get(getConfigEndpoint, getConfig(runtime.Config, runtime.Logger))
	r.Get(getUsersEndpoint, getUsers(runtime.UserRepository, runtime.Logger))
	r.Post(postUsersEndpoint, postUsers(runtime.UserRepository, runtime.Logger))

	return r
}
