package grpc

import (
	"github.com/espitman/gws-chat/message-service/internal/core/port"
	pb "github.com/espitman/gws-chat/pkg/protos/protogen/message-service"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	pb.UnimplementedMessageServiceServer
	validate       *validator.Validate
	roomService    port.RoomService
	memberService  port.MemberService
	messageService port.MessageService
	chatService    port.ChatService
	// +salvation Handler

}

func NewHandler(
	validate *validator.Validate,
	roomService port.RoomService,
	memberService port.MemberService,
	messageService port.MessageService,
	chatService port.ChatService,
	// +salvation NewHandlerType
) *Handler {
	return &Handler{
		validate:       validate,
		roomService:    roomService,
		memberService:  memberService,
		messageService: messageService,
		chatService:    chatService,
		// +salvation NewHandler
	}
}
