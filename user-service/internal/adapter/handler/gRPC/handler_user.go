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
			Token:    result.Token,
		},
	}
	return &resp, nil
}

func (h Handler) V1GetAll(ctx context.Context, req *pb.V1GetAllRequest) (*pb.V1GetAllResponse, error) {
	var resp pb.V1GetAllResponse
	var users []*pb.UserPublic
	result, err := h.userService.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, r := range result {
		users = append(users, &pb.UserPublic{
			Id:     int32(r.ID),
			Name:   r.Name,
			Avatar: r.Avatar,
			Status: r.Status,
		})
	}
	resp.Users = users

	return &resp, nil
}

func (h Handler) V1ValidateToken(ctx context.Context, req *pb.V1ValidateTokenRequest) (*pb.V1ValidateTokenResponse, error) {
	user, err := h.userService.ValidateToken(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	resp := pb.V1ValidateTokenResponse{
		User: &pb.UserPublic{
			Id:   int32(user.ID),
			Name: user.Name,
		},
	}
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (h Handler) V1Get(ctx context.Context, req *pb.V1GetRequest) (*pb.V1GetResponse, error) {
	result, err := h.userService.Get(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	return &pb.V1GetResponse{
		User: &pb.UserPublic{
			Id:     int32(result.ID),
			Name:   result.Name,
			Avatar: result.Avatar,
			Status: result.Status,
		},
	}, nil
}
