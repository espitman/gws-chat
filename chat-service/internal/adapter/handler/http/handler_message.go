package http

import (
	"github.com/espitman/gws-chat/chat-service/internal/core/port"
	"github.com/go-playground/validator/v10"
)

type MessageHandler struct {
	validate       *validator.Validate
	messageService port.MessageService
	socketService  port.SocketService
}

func NewMessageHandler(validate *validator.Validate, messageService port.MessageService, socketService port.SocketService) *MessageHandler {
	return &MessageHandler{
		validate:       validate,
		messageService: messageService,
		socketService:  socketService,
	}
}
