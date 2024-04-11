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
	indexHandler   *IndexHandler
	sseHandler     *SSEHandler
	// +salvation ServerType
}

func NewServer(
	port string,
	chatHandler *ChatHandler,
	messageHandler *MessageHandler,
	indexHandler *IndexHandler,
	sseHandler *SSEHandler,
	// +salvation NewServerType
) Server {
	return Server{
		port:           port,
		chatHandler:    chatHandler,
		messageHandler: messageHandler,
		indexHandler:   indexHandler,
		sseHandler:     sseHandler,
		// +salvation ServerHandler
	}
}

func (s *Server) Run() {

	router := httprouter.New()

	router.GET("/chat/:id", corsMiddleware(s.chatHandler.ChatHandler))
	router.GET("/event/:id", s.sseHandler.SSEHandler)

	router.GET("/", corsMiddleware(s.indexHandler.IndexHandler))
	//info()
	fmt.Println("Server start on port: 8085")
	log.Fatal(http.ListenAndServe(":8085", router))
}
