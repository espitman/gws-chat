package port

import (
	"context"

	"github.com/espitman/gws-chat/chat-service/internal/core/domain"
)

/**
 * messageService implemented by service.messageService interface
 */

type MessageService interface {
	Create(ctx context.Context, message domain.Message) (*domain.Message, error)
}
