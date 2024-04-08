package service

import (
	"context"

	grpcclientmessage "github.com/espitman/gws-chat/chat-service/internal/adapter/grpcclient/message"
	"github.com/espitman/gws-chat/chat-service/internal/core/domain"
)

/**
 * messageService implements port.messageService interface
 */

type MessageService struct {
	messageRepositoryGrpc grpcclientmessage.GrpcClientMessage
}

func NewMessageService(messageRepositoryGrpc grpcclientmessage.GrpcClientMessage) *MessageService {
	return &MessageService{
		messageRepositoryGrpc,
	}
}

func (s *MessageService) Create(ctx context.Context, message domain.Message) (*domain.Message, error) {
	return s.messageRepositoryGrpc.Create(ctx, message)
}
