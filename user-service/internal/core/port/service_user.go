package port

import (
	"context"

	"github.com/espitman/gws-chat/user-service/internal/core/domain"
)

/**
 * UserService implemented by service.UserService interface
 */

type UserService interface {
	Login(ctx context.Context, user domain.User) (*domain.User, error)
	GetAll(ctx context.Context) ([]*domain.User, error)
	ValidateToken(ctx context.Context, token string) (*domain.User, error)
	Get(ctx context.Context, userID uint32) (*domain.User, error)
}
