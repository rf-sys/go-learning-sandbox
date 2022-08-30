package routes

import (
	"awesomeProject1/log"
	"awesomeProject1/service"
	"encoding/json"
	"net/http"
)

const getUsersEndpoint = "/users"

func getUsers(service service.UserService, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := service.GetAllUsers()
		if err != nil {
			logger.Error(err, "failed finding users")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		usersJson, err := json.Marshal(users)
		if err != nil {
			logger.Error(err, "failed marshalling users")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(usersJson)
	}
}
