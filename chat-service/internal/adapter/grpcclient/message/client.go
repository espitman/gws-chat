package grpcclientmessage

import (
	"context"
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
		Body:   message.Body.Body,
		Time:   message.Time,
	}
	md := metadata.Pairs("Authorization", strconv.Itoa(int(message.UserID)))
	ctx = metadata.NewOutgoingContext(ctx, md)

	result, err := g.pb.V1AddMessage(ctx, &reqDto)
	if err != nil {
		return nil, err
	}
	return &domain.Message{
		ID:     result.Message.Id,
		RoomID: result.Message.RoomID,
		UserID: uint32(result.Message.UserID),
		Body: domain.MessageBody{
			Type: "message",
			Body: result.Message.Body,
		},
		Time: result.Message.Time,
	}, nil
}

func (g *GrpcClientMessage) GetAudienceID(ctx context.Context, roomID string, userID uint32) (uint32, error) {
	md := metadata.Pairs("Authorization", strconv.Itoa(int(userID)))
	ctx = metadata.NewOutgoingContext(ctx, md)
	reqDto := messagepb.V1GetAudienceIDRequest{RoomID: roomID}
	result, err := g.pb.V1GetAudienceID(ctx, &reqDto)
	if err != nil {
		return 0, err
	}
	return result.AudienceID, nil
}
