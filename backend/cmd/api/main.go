package main

import (
	"os"
	"president/internal/server"

	"github.com/charmbracelet/log"
)

func main() {
	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
	})
	logger.Info("starting server")

	server := server.NewServer(logger)

	err := server.ListenAndServe()
	if err != nil {
		log.Errorf("cannot start server: %s", err)
		panic("")
	}
}
