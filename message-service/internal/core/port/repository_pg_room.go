package port

import (
	"context"

	"github.com/espitman/gws-chat/message-service/internal/core/domain"
)

/**
 * roomRepository implemented by repository.roomRepository interface
 */

type RoomRepositoryPg interface {
	Crete(ctx context.Context, room domain.Room) (*domain.Room, bool, error)
	Get(ctx context.Context, roomID string) (*domain.Room, error)
	GetRooms(ctx context.Context, IDs []string) ([]*domain.Room, error)
}
