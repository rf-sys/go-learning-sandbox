package routes

import (
	"awesomeProject1/log"
	"awesomeProject1/model"
	"awesomeProject1/service"
	"encoding/json"
	"fmt"
	"net/http"
)

const postUsersEndpoint = "/users"

func postUsers(service service.UserService, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			logger.Error(err, "failed decoding JSON payload")
			http.Error(w, "Invalid payload", http.StatusBadRequest)
			return
		}

		newUser, err := service.Create(user)

		if err != nil {
			logger.Error(err, "failed inserting new user into database")
			http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
			return
		}

		w.WriteHeader(http.StatusCreated)
		_, err = w.Write([]byte(fmt.Sprintf("user created with id %v", newUser.ID)))
		if err != nil {
			logger.Error(err, "failed writing a response body")
		}
	}
}
