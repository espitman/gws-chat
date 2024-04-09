package port

import (
	"context"

	"github.com/espitman/gws-chat/message-service/internal/core/domain"
)

type UserRepositoryGrpc interface {
	Get(ctx context.Context, userID uint32) (*domain.User, error)
	GetByIds(ctx context.Context, userIDs []uint32) ([]*domain.User, error)
}
