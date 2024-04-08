package service

import (
	"context"

	"github.com/espitman/gws-chat/message-service/internal/core/domain"
	"github.com/espitman/gws-chat/message-service/internal/core/port"
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
