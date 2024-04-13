package api

import (
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	superConf "github.com/espitman/go-super-conf"
	"github.com/espitman/gws-chat/chat-service/internal/adapter/handler/http"
	"github.com/espitman/gws-chat/chat-service/internal/core/port"
	"github.com/go-playground/validator/v10"
	"github.com/r3labs/sse/v2"
)

func Run(
	validate *validator.Validate,
	pubSub *gochannel.GoChannel,
	messageService port.MessageService,
	socketConnectService port.SocketConnetService,
	socketService port.SocketService,
	roomService port.RoomService,
	userService port.UserService,
	// +salvation RunType
) {

	sseServer := sse.New()

	chatHandler := http.NewChatHandler(validate, messageService, socketConnectService, socketService, roomService, userService)
	messageHandler := http.NewMessageHandler(validate, messageService, socketService)
	indexHandler := http.NewIndexHandler()
	sseHandler := http.NewSSEHandler(sseServer, pubSub)
	// +salvation NewHandler

	httpServer := http.NewServer(
		superConf.Get("app.http.port"),
		chatHandler,
		messageHandler,
		indexHandler,
		sseHandler,
		// +salvation NewServerHandler
	)

	httpServer.Run()
}
