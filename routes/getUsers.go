package routes

import (
	"awesomeProject1/model"
	"awesomeProject1/repository"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

func postUsers(repository repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Err(err).Msgf("failed decoding JSON payload: %v", err)
			http.Error(w, "Invalid payload", http.StatusBadRequest)
			return
		}

		err = repository.Create(user)
		if err != nil {
			log.Err(err).Msgf("failed inserting new user into database: %v", err)
			http.Error(w, "Username already exists", http.StatusUnprocessableEntity)
			return
		}

		w.WriteHeader(http.StatusCreated)
		_, err = w.Write([]byte("user created"))
		if err != nil {
			log.Err(err).Msg("failed writing a response body")
		}
	}
}
