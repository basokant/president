package server

import (
	"encoding/json"
	"net/http"

	"github.com/charmbracelet/log"
)

func (s *Server) RegisterRoutes(logger *log.Logger) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", s.healthHandler(logger))

	return mux
}

func (s *Server) healthHandler(logger *log.Logger) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logRequest(logger, r)

		jsonResp, err := json.Marshal(s.db.Health())
		if err != nil {
			logger.Error("error handling JSON marshal. Err: %v", err)
		}

		_, _ = w.Write(jsonResp)
	})
}
