package port

import (
	"context"

	"github.com/espitman/gws-chat/chat-service/internal/core/domain"
)

type UserService interface {
	Validate(ctx context.Context, token string) (*domain.User, error)
	SetOnline(ctx context.Context, userID uint32, isOnline bool) (*domain.User, error)
}
