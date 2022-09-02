package handlers

import (
	"awesomeProject1/log"
	"awesomeProject1/service"
	"encoding/json"
	"net/http"
)

type GetUsersHandler struct {
	service service.UserService
	logger  log.Logger
}

func NewGetUsersHandler(service service.UserService, logger log.Logger) GetUsersHandler {
	return GetUsersHandler{service: service, logger: logger}
}

func (h GetUsersHandler) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := h.service.GetAllUsers()
		if err != nil {
			h.logger.Error(err, "failed finding users")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		usersJson, err := json.Marshal(users)
		if err != nil {
			h.logger.Error(err, "failed marshalling users")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(usersJson)
	}
}
