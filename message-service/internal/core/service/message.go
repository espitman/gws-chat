package service

import (
	"context"

	"github.com/espitman/gws-chat/message-service/internal/core/domain"
	"github.com/espitman/gws-chat/message-service/internal/core/port"
	messagepb "github.com/espitman/gws-chat/pkg/protos/protogen/message-service"
)

/**
 * MessageService implements port.MessageService interface
 */

type MessageService struct {
	messageRepositoryPg port.MessageRepositoryPg
}

func NewMessageService(
	messageRepositoryPg port.MessageRepositoryPg,
) *MessageService {
	return &MessageService{
		messageRepositoryPg: messageRepositoryPg,
	}
}

func (s *MessageService) Crete(ctx context.Context, message domain.Message) (*domain.Message, error) {
	pgMessage, err := s.messageRepositoryPg.Crete(ctx, message)
	if err != nil {
		return nil, err
	}
	return &domain.Message{
		ID:     pgMessage.ID,
		RoomID: pgMessage.RoomID,
		UserID: pgMessage.UserID,
		Body:   pgMessage.Body,
		Time:   pgMessage.Time,
	}, nil
}

func (s *MessageService) Get(ctx context.Context, ID int) (*domain.Message, error) {
	pgMessage, err := s.messageRepositoryPg.Get(ctx, ID)
	if err != nil {
		return nil, err
	}
	return &domain.Message{
		ID:     pgMessage.ID,
		RoomID: pgMessage.RoomID,
		UserID: pgMessage.UserID,
		Body:   pgMessage.Body,
		Time:   pgMessage.Time,
	}, nil
}

func (s *MessageService) GetRoomMessages(ctx context.Context, roomInfo *domain.RoomInfo) ([]*domain.Message, error) {
	messages, err := s.messageRepositoryPg.GetRoomMessages(ctx, roomInfo.RoomID)
	if err != nil {
		return nil, err
	}
	for i, message := range messages {
		if message.UserID == roomInfo.User.ID {
			messages[i].Type = messagepb.MessageType_TYPE_SENT
			messages[i].UserName = roomInfo.User.Name
			messages[i].UserAvatar = roomInfo.User.Avatar
		} else {
			messages[i].Type = messagepb.MessageType_TYPE_RECEIVED
			messages[i].UserName = roomInfo.Audience.Name
			messages[i].UserAvatar = roomInfo.Audience.Avatar
		}

	}
	return messages, nil
}
