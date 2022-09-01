package handlers

import (
	"awesomeProject1/config"
	"awesomeProject1/log"
	"encoding/json"
	"net/http"
)

type GetConfigHandler struct {
	cfg    config.Config
	logger log.Logger
}

func (h GetConfigHandler) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cfgJson, err := json.Marshal(h.cfg)
		if err != nil {
			h.logger.Error(err, "failed marshalling configs")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write(cfgJson)
	}
}

func NewGetConfigHandler(cfg config.Config, logger log.Logger) GetConfigHandler {
	return GetConfigHandler{
		cfg:    cfg,
		logger: logger,
	}
}
