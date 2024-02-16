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

	mux.HandleFunc("GET /chat-room", s.getChatHandler())
	mux.HandleFunc("GET /chat-stream", s.getChatStreamHandler(logger))
	mux.HandleFunc("POST /chat", s.postChatHandler(logger))

	return mux
}

func (s *Server) healthHandler(logger *log.Logger) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logRequest(logger, r)

		// Set CORS headers to allow all origins. You may want to restrict this to specific origins in a production environment.
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, PUT, GET, OPTIONS")
		w.Header().Set("Access-Control-Expose-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")

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
		w.Header().Set("Access-Control-Allow-Methods", "POST, PUT, GET, OPTIONS")
		w.Header().Set("Access-Control-Expose-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")

		sse.WriteHeaders(w)
		// Simulate sending events (you can replace this with real data)
		for i := 0; i < 10; i++ {
			e := sse.Event[int]{
				Data: i,
			}
			sse.Encode[int](w, e)

			time.Sleep(2 * time.Second)
			w.(http.Flusher).Flush()
		}
	})
}

func (s *Server) getChatHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		html.Execute(w, struct{}{})
	})
}

func (s *Server) getChatStreamHandler(logger *log.Logger) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logRequest(logger, r)

		go func() {
			// Received Browser Disconnection
			<-r.Context().Done()
			println("The client is disconnected here")
		}()

		// Set CORS headers to allow all origins. You may want to restrict this to specific origins in a production environment.
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, PUT, GET, OPTIONS")
		w.Header().Set("Access-Control-Expose-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")

		sse.WriteHeaders(w)

		ch := s.db.SubscribeToChannel("chat")

		for msg := range ch {
			if msg != nil && msg.Payload != "" {
				logger.Print("received from chat", "message", msg.Payload)
				e := sse.Event[string]{
					Data: msg.Payload,
				}
				sse.Encode(w, e)
				w.(http.Flusher).Flush()
			}
		}

		logger.Print("closing chat stream")
	})
}

type ChatMessage struct {
	Message string `json:"message"`
}

func (s *Server) postChatHandler(logger *log.Logger) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logRequest(logger, r)

		// Set CORS headers to allow all origins. You may want to restrict this to specific origins in a production environment.
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, PUT, GET, OPTIONS")
		w.Header().Set("Access-Control-Expose-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")

		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			logger.Error("error parsing form. Err: %v", err)
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		}

		message := r.FormValue("message")

		err = s.db.SendMessage("chat", message)
		if err != nil {
			logger.Error("error sending message. Err: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		logger.Info("Message sent", "message", message)
		fmt.Fprint(w, "success")
	})
}
