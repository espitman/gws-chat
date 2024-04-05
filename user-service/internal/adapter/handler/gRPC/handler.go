package grpc

import (
	pb "github.com/espitman/gws-chat/pkg/protos/protogen/user-service"
	"github.com/espitman/gws-chat/user-service/internal/core/port"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	pb.UnimplementedUserServiceServer
	validate    *validator.Validate
	userService port.UserService
	// +salvation Handler
}

func NewHandler(
	validate *validator.Validate,
	userService port.UserService,
	// +salvation NewHandlerType
) *Handler {
	return &Handler{
		validate:    validate,
		userService: userService,
		// +salvation NewHandlerService
	}
}
