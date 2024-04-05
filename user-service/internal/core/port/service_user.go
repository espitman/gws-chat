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
}
