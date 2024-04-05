package grpc

import (
	"context"
	pb "github.com/espitman/gws-chat/pkg/protos/protogen/user-service"
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

func (h Handler) V1GetAll(ctx context.Context, req *pb.V1GetAllRequest) (*pb.V1GetAllResponse, error) {
	var resp pb.V1GetAllResponse
	var users []*pb.User
	result, err := h.userService.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, r := range result {
		users = append(users, &pb.User{
			Id:     int32(r.ID),
			Name:   r.Name,
			Avatar: r.Avatar,
			Status: r.Password,
		})
	}
	resp.Users = users

	return &resp, nil
}
