package port

import (
	"context"
	"github.com/espitman/gws-chat/user-service/internal/core/domain"
)

/**
 * userRepository implemented by repository.userRepository interface
 */

type UserRepositoryPg interface {
	Crete(ctx context.Context, user domain.User) (*domain.User, error)
	Get(ctx context.Context, user domain.User) (*domain.User, error)
	GetAll(ctx context.Context) ([]*domain.User, error)
}
