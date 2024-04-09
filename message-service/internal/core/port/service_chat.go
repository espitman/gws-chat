package port

import (
	"context"

	"github.com/espitman/gws-chat/message-service/internal/core/domain"
)

/**
 * ChatService implemented by service.ChatService interface
 */

type ChatService interface {
	GetUserChats(ctx context.Context, userID uint32) ([]*domain.Chat, error)
}
