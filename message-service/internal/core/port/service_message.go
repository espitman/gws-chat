package port

import (
	"context"

	"github.com/espitman/gws-chat/message-service/internal/core/domain"
)

/**
 * MessageService implemented by service.MessageService interface
 */

type MessageService interface {
	Crete(ctx context.Context, message domain.Message) (*domain.Message, error)
	Get(ctx context.Context, ID int) (*domain.Message, error)
}
