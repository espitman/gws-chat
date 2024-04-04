package port

import (
	"context"
	"github.com/espitman/gws-chat/chat-service/internal/core/domain"
)

/**
 * messageService implemented by service.messageService interface
 */

type MessageService interface {
	Crete(ctx context.Context, message domain.Message) (*domain.Message, error)
	Get(ctx context.Context, ID int) (*domain.Message, error)
	Update(ctx context.Context, ID int, message domain.Message) (*domain.Message, error)
	Delete(ctx context.Context, ID int) (*domain.Message, error)
}
