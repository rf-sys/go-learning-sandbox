package handlers

import (
	"awesomeProject1/log"
	"awesomeProject1/model"
	"awesomeProject1/service"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type EditUserHandler struct {
	service service.UserService
	logger  log.Logger
}

func NewEditUserHandler(service service.UserService, logger log.Logger) EditUserHandler {
	return EditUserHandler{service: service, logger: logger}
}

func (h EditUserHandler) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User

		param := chi.URLParam(r, "id")
		if param == "" {
			http.Error(w, "no user id", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(param)
		if err != nil {
			http.Error(w, "invalid user id", http.StatusBadRequest)
			return
		}

		err = json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			h.logger.Error(err, "failed decoding JSON payload")
			http.Error(w, "invalid payload", http.StatusBadRequest)
			return
		}

		err = h.service.Edit(id, user)
		if err != nil {
			h.logger.Error(err, "failed updating the user")

			if errors.Is(err, service.ErrUserNotFound) {
				http.Error(w, err.Error(), http.StatusUnprocessableEntity)
				return
			}

			http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(fmt.Sprintf("user updated with id %v", id)))
	}
}
