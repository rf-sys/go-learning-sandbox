package routes

import (
	"awesomeProject1/runtime"
	"github.com/go-chi/chi/v5"
)

const getUsersEndpoint = "/users"

func LoadRoutes(r *chi.Mux) *chi.Mux {
	r.Get(getConfigEndpoint, getConfig(runtime.Config))
	r.Get(getUsersEndpoint, getUsers(runtime.UserRepository))
	r.Post(postUsersEndpoint, postUsers(runtime.UserRepository))

	return r
}
