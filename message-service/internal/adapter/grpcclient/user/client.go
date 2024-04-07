package grpcclientuser

import (
	"context"
	"fmt"

	"github.com/espitman/gws-chat/message-service/internal/core/domain"
	userpb "github.com/espitman/gws-chat/pkg/protos/protogen/user-service"
	"google.golang.org/grpc"
)

type GrpcClientUser struct {
	pb userpb.UserServiceClient
}

func NewGrpcClientUser() *GrpcClientUser {
	cc, err := grpc.Dial("localhost:8001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	pb := userpb.NewUserServiceClient(cc)
	return &GrpcClientUser{
		pb: pb,
	}
}

func (g *GrpcClientUser) Get(ctx context.Context, userID uint32) (*domain.User, error) {
	reqDto := userpb.V1GetRequest{
		UserID: userID,
	}
	result, err := g.pb.V1Get(ctx, &reqDto)
	fmt.Println(result)
	if err != nil {
		return nil, err
	}
	return &domain.User{
		ID:     uint32(result.User.Id),
		Name:   result.User.Name,
		Avatar: result.User.Avatar,
		Status: result.User.Status,
	}, nil
}
