package service

import (
	"context"
	"github.com/espitman/gws-chat/chat-service/internal/core/domain"
	"github.com/espitman/gws-chat/chat-service/internal/core/port"
)

/**
 * messageService implements port.messageService interface
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
	var result domain.Message
	pgMessage, err := s.messageRepositoryPg.Crete(ctx, message)
	if err != nil {
		return nil, err
	}
	result.ID = pgMessage.ID
	result.Name = pgMessage.Name
	message.ID = pgMessage.ID
	return &result, nil
}

func (s *MessageService) Get(ctx context.Context, ID int) (*domain.Message, error) {
	var result domain.Message
	pgMessage, err := s.messageRepositoryPg.Get(ctx, ID)
	if err != nil {
		return nil, err
	}
	result.ID = pgMessage.ID
	result.Name = pgMessage.Name
	return &result, nil
}

func (s *MessageService) Update(ctx context.Context, ID int, message domain.Message) (*domain.Message, error) {
	var result domain.Message
	pgMessage, err := s.messageRepositoryPg.Update(ctx, ID, message)
	if err != nil {
		return nil, err
	}
	result = *pgMessage
	return &result, nil
}

func (s *MessageService) Delete(ctx context.Context, ID int) (*domain.Message, error) {
	var result domain.Message
	var err error
	pgMessage, err := s.messageRepositoryPg.Delete(ctx, ID)
	if err != nil {
		return nil, err
	}
	result = *pgMessage
	return &result, nil
}
