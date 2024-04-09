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
	GetByName(ctx context.Context, user domain.User) (*domain.User, error)
	GetAll(ctx context.Context, me bool, userID int) ([]*domain.User, error)
	Get(ctx context.Context, userID uint32) (*domain.User, error)
	GetByIDs(ctx context.Context, userIDs []int) ([]*domain.User, error)
}
