package routes

import (
	"awesomeProject1/repository"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

const postUsersEndpoint = "/users"

func getUsers(repository repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := repository.FindAll()
		if err != nil {
			log.Err(err).Msgf("failed finding users")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		usersJson, err := json.Marshal(users)
		if err != nil {
			log.Err(err).Msg("failed marshalling users")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(usersJson)
	}
}
