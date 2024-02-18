package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"president/internal/database"

	"github.com/charmbracelet/log"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	db     database.Service
	logger *log.Logger
	port   int
}

func NewServer(logger *log.Logger) *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		db:     database.New(),
		logger: logger,
		port:   port,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

func (s Server) NewRequestLogger(r *http.Request) *log.Logger {
	s.logger.Helper()
	message := fmt.Sprintf("%s %s", r.Method, r.URL.Path)
	s.logger.Info(message)
	return s.logger.With("method", r.Method, "URL", r.URL.Path)
}
