package grpcclientuser

import (
	"context"

	"github.com/espitman/gws-chat/message-service/internal/core/domain"
	userpb "github.com/espitman/gws-chat/pkg/protos/protogen/user-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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
	myUserIDCtx := ctx.Value("userID")
	myUserID := myUserIDCtx.(string)
	md := metadata.Pairs("Authorization", myUserID)
	ctx = metadata.NewOutgoingContext(ctx, md)

	reqDto := userpb.V1GetRequest{
		UserID: userID,
	}
	result, err := g.pb.V1Get(ctx, &reqDto)
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

func (g *GrpcClientUser) GetByIds(ctx context.Context, userIDs []uint32) ([]*domain.User, error) {
	myUserIDCtx := ctx.Value("userID")
	myUserID := myUserIDCtx.(string)
	md := metadata.Pairs("Authorization", myUserID)
	ctx = metadata.NewOutgoingContext(ctx, md)
	reqDto := userpb.V1GetByIDsRequest{
		UserIDs: userIDs,
	}
	result, err := g.pb.V1GetByIDs(ctx, &reqDto)
	if err != nil {
		return nil, err
	}
	var users []*domain.User
	for _, user := range result.Users {
		users = append(users, &domain.User{
			ID:     uint32(user.Id),
			Name:   user.Name,
			Avatar: user.Avatar,
			Status: user.Status,
		})
	}
	return users, nil
}
