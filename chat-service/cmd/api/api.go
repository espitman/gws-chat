package api

import (
	superConf "github.com/espitman/go-super-conf"
	"github.com/espitman/gws-chat/chat-service/internal/adapter/handler/http"
	"github.com/espitman/gws-chat/chat-service/internal/core/port"
	"github.com/go-playground/validator/v10"
)

func Run(
	validate *validator.Validate,
	messageService port.MessageService,
	socketConnectService port.SocketConnetService,
	socketService port.SocketService,
	roomService port.RoomService,
	userService port.UserService,
	// +salvation RunType
) {

	chatHandler := http.NewChatHandler(validate, messageService, socketConnectService, socketService, roomService, userService)
	messageHandler := http.NewMessageHandler(validate, messageService, socketService)
	// +salvation NewHandler

	httpServer := http.NewServer(
		superConf.Get("app.http.port"),
		chatHandler,
		messageHandler,
		// +salvation NewServerHandler
	)
	httpServer.Run()
}
