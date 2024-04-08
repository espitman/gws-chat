package port

import (
	"context"

	"github.com/espitman/gws-chat/chat-service/internal/core/domain"
)

type UserService interface {
	Validate(ctx context.Context, token string) (*domain.User, error)
}
