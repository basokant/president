package server

import (
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
)

func logRequest(logger *log.Logger, r *http.Request) {
	logger.Helper()
	message := fmt.Sprintf("%s %s", r.Method, r.URL.Path)
	logger.Info(message)
}
