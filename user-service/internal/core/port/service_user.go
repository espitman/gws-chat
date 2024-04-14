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
	GetAll(ctx context.Context, me bool, userID uint32) ([]*domain.User, error)
	ValidateToken(ctx context.Context, token string) (*domain.User, error)
	Get(ctx context.Context, userID uint32) (*domain.User, error)
	GetByIDs(ctx context.Context, userIDs []uint32) ([]*domain.User, error)
	SetOnline(ctx context.Context, userID uint32, isOnline bool) (*domain.User, error)
}
