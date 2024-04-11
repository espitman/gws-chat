package port

import (
	"context"

	"github.com/espitman/gws-chat/chat-service/internal/core/domain"
)

type MessageRepositoryGrpc interface {
	Create(ctx context.Context, message domain.Message) (*domain.Message, error)
	GetAudienceID(ctx context.Context, roomID string, userID uint32) (uint32, error)
}
