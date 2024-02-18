package server

import "net/http"

// Set CORS headers to allow all origins. You may want to restrict this to specific origins in a production environment.
func writeCorsHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, PUT, GET, OPTIONS")
	w.Header().Set("Access-Control-Expose-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
}
