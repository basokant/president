package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"president/internal/sse"
	"time"

	"github.com/charmbracelet/log"
)

func (s *Server) RegisterRoutes(logger *log.Logger) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", s.healthHandler(logger))
	mux.HandleFunc("GET /events", s.eventsHandler(logger))

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

func (s *Server) eventsHandler(logger *log.Logger) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logRequest(logger, r)

		// Set CORS headers to allow all origins. You may want to restrict this to specific origins in a production environment.
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Type")

		sse.WriteHeaders(w)
		// Simulate sending events (you can replace this with real data)
		for i := 0; i < 10; i++ {
			fmt.Fprintf(w, "data: %s\n\n", fmt.Sprintf("Event %d", i))
			time.Sleep(2 * time.Second)
			w.(http.Flusher).Flush()
		}
	})
}
