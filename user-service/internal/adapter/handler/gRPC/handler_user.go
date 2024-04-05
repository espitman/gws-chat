package grpc

import (
	"context"
	pb "github.com/espitman/gws-chat/pkg/protos/protogen/order-service"
	"github.com/espitman/gws-chat/user-service/internal/core/domain"
)

func (h Handler) V1Login(ctx context.Context, req *pb.V1LoginRequest) (*pb.V1LoginResponse, error) {
	user := domain.User{
		Name:     req.Name,
		Password: req.Password,
	}
	result, err := h.userService.Login(ctx, user)
	if err != nil {
		return nil, err
	}

	resp := pb.V1LoginResponse{
		User: &pb.User{
			Id:       int32(result.ID),
			Name:     result.Name,
			Password: result.Password,
			Avatar:   result.Avatar,
			Status:   result.Status,
		},
	}
	return &resp, nil
}
