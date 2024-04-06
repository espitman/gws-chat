package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Server struct {
	port           string
	chatHandler    *ChatHandler
	messageHandler *MessageHandler
	// +salvation ServerType
}

func NewServer(
	port string,
	chatHandler *ChatHandler,
	messageHandler *MessageHandler,
	// +salvation NewServerType
) Server {
	return Server{
		port:           port,
		chatHandler:    chatHandler,
		messageHandler: messageHandler,
		// +salvation ServerHandler
	}
}

func (s *Server) Run() {
	router := httprouter.New()

	router.GET("/chat/:id", corsMiddleware(s.chatHandler.ChatHandler))
	router.GET("/", corsMiddleware(Index))

	//info()
	log.Fatal(http.ListenAndServe(":8085", router))

}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, _ = fmt.Fprint(w, "Welcome!\n")
}

func corsMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		header := w.Header()
		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Access-Control-Allow-Credentials", "true")
		header.Set("Access-Control-Allow-Headers", "*")
		header.Set("Access-Control-Allow-Methods", "*")
		next(w, r, ps)
	}
}
