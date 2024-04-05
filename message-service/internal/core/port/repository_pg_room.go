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
}
