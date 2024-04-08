package grpcclientmessage

import (
	"context"
	"fmt"
	"strconv"

	"github.com/espitman/gws-chat/chat-service/internal/core/domain"
	messagepb "github.com/espitman/gws-chat/pkg/protos/protogen/message-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type GrpcClientMessage struct {
	pb messagepb.MessageServiceClient
}

func NewGrpcClientMessage() *GrpcClientMessage {
	cc, err := grpc.Dial("localhost:8002", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	pb := messagepb.NewMessageServiceClient(cc)
	return &GrpcClientMessage{
		pb: pb,
	}
}

func (g *GrpcClientMessage) Create(ctx context.Context, message domain.Message) (*domain.Message, error) {
	reqDto := messagepb.V1AddMessageRequest{
		RoomID: message.RoomID,
		Body:   message.Body,
		Time:   message.Time,
	}
	fmt.Println("@@@@", &reqDto)

	md := metadata.Pairs("Authorization", strconv.Itoa(int(message.UserID)))
	ctx = metadata.NewOutgoingContext(ctx, md)

	result, err := g.pb.V1AddMessage(ctx, &reqDto)
	if err != nil {
		return nil, err
	}
	return &domain.Message{
		ID:     result.Id,
		RoomID: result.RoomID,
		UserID: uint32(result.UserId),
		Body:   result.Body,
		Time:   result.Time,
	}, nil
}
