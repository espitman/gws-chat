package port

import (
	"context"

	"github.com/espitman/gws-chat/chat-service/internal/core/domain"
	"github.com/lxzan/gws"
)

/**
 * messageService implemented by service.messageService interface
 */

type MessageService interface {
	Create(ctx context.Context, message domain.Message) (*domain.Message, error)
	Text(ctx context.Context, message *gws.Message, socketID string, roomID string, userID uint32, messageBody domain.MessageBody)
}
