package grpc

import (
	"github.com/espitman/gws-chat/message-service/internal/core/port"
	pb "github.com/espitman/gws-chat/pkg/protos/protogen/message-service"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	pb.UnimplementedMessageServiceServer
	validate      *validator.Validate
	roomService   port.RoomService
	memberService port.MemberService
	// +salvation Handler

}

func NewHandler(
	validate *validator.Validate,
	roomService port.RoomService,
	memberService port.MemberService,
	// +salvation NewHandlerType
) *Handler {
	return &Handler{
		validate:      validate,
		roomService:   roomService,
		memberService: memberService,
		// +salvation NewHandler
	}
}
