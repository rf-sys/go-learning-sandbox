package routes

import (
	"awesomeProject1/config"
	"awesomeProject1/log"
	"encoding/json"
	"net/http"
)

const getConfigEndpoint = "/configs"

func getConfig(cfg config.Config, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cfgJson, err := json.Marshal(cfg)
		if err != nil {
			logger.Error(err, "failed marshalling configs")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write(cfgJson)
	}
}
