package port

import (
	"context"

	"github.com/espitman/gws-chat/message-service/internal/core/domain"
)

type UserService interface {
	Get(ctx context.Context, userID uint32) (*domain.User, error)
}
