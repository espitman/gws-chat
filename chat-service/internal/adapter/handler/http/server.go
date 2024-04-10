package http

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Server struct {
	port           string
	chatHandler    *ChatHandler
	messageHandler *MessageHandler
	indexHandler   *IndexHandler
	// +salvation ServerType
}

func NewServer(
	port string,
	chatHandler *ChatHandler,
	messageHandler *MessageHandler,
	indexHandler *IndexHandler,
	// +salvation NewServerType
) Server {
	return Server{
		port:           port,
		chatHandler:    chatHandler,
		messageHandler: messageHandler,
		indexHandler:   indexHandler,
		// +salvation ServerHandler
	}
}

func (s *Server) Run() {
	router := httprouter.New()

	router.GET("/chat/:id", corsMiddleware(s.chatHandler.ChatHandler))
	router.GET("/", corsMiddleware(s.indexHandler.IndexHandler))

	//info()
	log.Fatal(http.ListenAndServe(":8085", router))

}
