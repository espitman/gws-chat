package port

import (
	"context"
	"github.com/espitman/gws-chat/message-service/internal/core/domain"
)

/**
 * roomRepository implemented by repository.roomRepository interface
 */

type RoomRepositoryPg interface {
	Crete(ctx context.Context, room domain.Room) (*domain.Room, error)
	Get(ctx context.Context, ID int) (*domain.Room, error)
	Update(ctx context.Context, ID int, room domain.Room) (*domain.Room, error)
	Delete(ctx context.Context, ID int) (*domain.Room, error)
}
