package gRPC

import (
	superConf "github.com/espitman/go-super-conf"
	grpc "github.com/espitman/gws-chat/message-service/internal/adapter/handler/gRPC"
	"github.com/espitman/gws-chat/message-service/internal/core/port"
	"github.com/go-playground/validator/v10"
)

func Run(
	validate *validator.Validate,
	roomService port.RoomService,
	memberService port.MemberService,
	messageService port.MessageService,
	// +salvation Run
) {
	gRPC := grpc.NewServer(
		superConf.Get("app.grpc.port"),
		validate,
		roomService,
		memberService,
		messageService,
		// +salvation RunService
	)
	gRPC.Run()
}
