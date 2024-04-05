package gRPC

import (
	superConf "github.com/espitman/go-super-conf"
	grpc "github.com/espitman/gws-chat/user-service/internal/adapter/handler/gRPC"
	"github.com/espitman/gws-chat/user-service/internal/core/port"
	"github.com/go-playground/validator/v10"
)

func Run(
	validate *validator.Validate,
	userService port.UserService,
	// +salvation Run
) {
	gRPC := grpc.NewServer(
		superConf.Get("app.grpc.port"),
		validate,
		userService,
		// +salvation RunService
	)
	gRPC.Run()
}
