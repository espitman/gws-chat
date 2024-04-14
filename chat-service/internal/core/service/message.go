package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	grpcclientmessage "github.com/espitman/gws-chat/chat-service/internal/adapter/grpcclient/message"
	"github.com/espitman/gws-chat/chat-service/internal/core/domain"
	"github.com/espitman/gws-chat/chat-service/internal/core/port"
	"github.com/lxzan/gws"
)

/**
 * messageService implements port.messageService interface
 */

type MessageService struct {
	pubSub                *gochannel.GoChannel
	messageRepositoryGrpc grpcclientmessage.GrpcClientMessage
	roomService           port.RoomService
}

func NewMessageService(
	pubSub *gochannel.GoChannel,
	messageRepositoryGrpc grpcclientmessage.GrpcClientMessage,
	roomService port.RoomService,
) *MessageService {
	return &MessageService{
		pubSub,
		messageRepositoryGrpc,
		roomService,
	}
}

func (s *MessageService) Create(ctx context.Context, message domain.Message) (*domain.Message, error) {
	return s.messageRepositoryGrpc.Create(ctx, message)
}

func (s *MessageService) Text(ctx context.Context, message *gws.Message, socketID string, roomID string, userID uint32, messageBody domain.MessageBody) {
	audienceID, err := s.roomService.GetAudience(ctx, roomID, userID)
	if err != nil {
		fmt.Println(err)
	}

	msg := domain.Message{
		RoomID:     roomID,
		UserID:     userID,
		Body:       messageBody,
		AudienceID: audienceID,
	}
	_, err = s.Create(context.Background(), msg)
	if err != nil {
		fmt.Println(err)
	}

	go s.publishMessages(strconv.Itoa(int(audienceID)), msg)

	subscribers := s.roomService.GetSubscribers(roomID)
	for _, sc := range subscribers {
		sid, _ := sc.Session().Load("socketID")
		if sid != socketID {
			sc.WriteMessage(message.Opcode, message.Bytes())
		}
	}
}

func (s *MessageService) publishMessages(userID string, data domain.Message) {
	go s.publishSSE(userID, data)
	go s.publishPush(userID, data)
}

func (s *MessageService) publishSSE(userID string, data domain.Message) {
	msg := message.NewMessage(watermill.NewUUID(), message.Payload(data.Body.Body))
	if err := s.pubSub.Publish(userID, msg); err != nil {
		fmt.Println(err)
	}
}

func (s *MessageService) publishPush(userID string, data domain.Message) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	msg := message.NewMessage(watermill.NewUUID(), jsonData)
	if err := s.pubSub.Publish("chat-message", msg); err != nil {
		fmt.Println(err)
	}
}
