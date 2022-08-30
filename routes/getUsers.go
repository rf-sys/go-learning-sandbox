package routes

import (
	"awesomeProject1/log"
	"awesomeProject1/model"
	"awesomeProject1/repository"
	"encoding/json"
	"fmt"
	"net/http"
)

const postUsersEndpoint = "/users"

func postUsers(repository repository.UserRepository, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			logger.Error(err, fmt.Sprintf("failed decoding JSON payload: %v", err))
			http.Error(w, "Invalid payload", http.StatusBadRequest)
			return
		}

		err = repository.Create(user)
		if err != nil {
			logger.Error(err, fmt.Sprintf("failed inserting new user into database: %v", err))
			http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
			return
		}

		w.WriteHeader(http.StatusCreated)
		_, err = w.Write([]byte("user created"))
		if err != nil {
			logger.Error(err, "failed writing a response body")
		}
	}
}
