package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"president/internal/database"
	"president/internal/sse"
	"strconv"
	"time"

	"github.com/charmbracelet/log"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", s.getHealthHandler())

	mux.HandleFunc("GET /events", s.streamEventsHandler())

	mux.HandleFunc("GET /chat-room", s.getChatRoomHandler())
	mux.HandleFunc("GET /chat-stream", s.streamChatHandler())
	mux.HandleFunc("POST /chat", s.postChatHandler())

	return mux
}

func (s *Server) getHealthHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.NewRequestLogger(r)

		writeCorsHeaders(w)

		jsonResp, err := json.Marshal(s.db.Health())
		if err != nil {
			logger.Error("error handling JSON marshal. Err: %v", err)
		}

		_, _ = w.Write(jsonResp)
	})
}

func (s *Server) streamEventsHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = s.NewRequestLogger(r)

		writeCorsHeaders(w)
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

func (s *Server) getChatRoomHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		html.Execute(w, struct{}{})
	})
}

func (s *Server) streamChatHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.NewRequestLogger(r)

		go func() {
			<-r.Context().Done()
			logger.Info("The client is disconnected here")
		}()

		writeCorsHeaders(w)
		sse.WriteHeaders(w)

		lastEventIdStr := w.Header().Get("Last-Event-Id")
		lastEventId, err := strconv.ParseInt(lastEventIdStr, 10, 64)

		if err != nil {
			lastEventId = 0
		}

		logger.Info("Get previous messages from", "event-id", lastEventId, "id string", lastEventIdStr)
		messages := s.db.GetPreviousMessagesFrom("chat", lastEventId)
		for _, m := range messages {
			id := uint(m["id"].(float64))
			text := m["text"].(string)
			msg := database.Message{
				Id:   id,
				Text: text,
			}

			sendEvent(w, msg)
		}
		w.(http.Flusher).Flush()

		ch := s.db.SubscribeToChannel("chat")

		for m := range ch {
			if m == nil || m.Payload == "" {
				continue
			}

			var msg database.Message
			json.Unmarshal([]byte(m.Payload), &msg)

			logger.Info("message received from chat", "text", msg.Text, "id", msg.Id)
			sendEvent(w, msg)
			w.(http.Flusher).Flush()
		}

		logger.Info("closing chat stream")
	})
}

func sendEvent(w http.ResponseWriter, msg database.Message) {
	e := sse.Event[string]{
		Id:    fmt.Sprint(msg.Id),
		Data:  msg.Text,
		Retry: 50,
	}
	log.Debug("sending event", "id", e.Id, "data", e.Data)

	sse.Encode(w, e)
}

func (s *Server) postChatHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.NewRequestLogger(r)

		writeCorsHeaders(w)

		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			logger.Errorf("error parsing form. Err: %v", err)
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		}

		msgText := r.FormValue("message")
		msg := database.NewMessage(msgText)

		err = s.db.SendMessage("chat", msg)
		if err != nil {
			logger.Errorf("error sending message. Err: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		logger.Info("Message sent", "message", msgText)
		fmt.Fprint(w, "success")
	})
}
