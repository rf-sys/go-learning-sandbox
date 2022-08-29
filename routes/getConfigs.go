package routes

import (
	"awesomeProject1/config"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

const getConfigEndpoint = "/configs"

func getConfig(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cfgJson, err := json.Marshal(cfg)
		if err != nil {
			log.Err(err).Msgf("failed marshalling configs")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write(cfgJson)
	}
}
