package http

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
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
	//app := fiber.New()
	//routes := newRouter(
	//    s.messageHandler,
	//    // +salvation RouterRun
	//)
	//routes.serve(app)
	//_ = app.Listen(":" + s.port)

	router := httprouter.New()
	router.GET("/chat/:id", s.chatHandler.ChatHandler)
	fmt.Println("App run on port 6666")
	//info()
	log.Fatal(http.ListenAndServe(":6666", router))

}
