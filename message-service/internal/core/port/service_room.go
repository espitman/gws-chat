package port

import (
	"context"

	"github.com/espitman/gws-chat/message-service/internal/core/domain"
)

/**
 * RoomService implemented by service.RoomService interface
 */

type RoomService interface {
	Crete(ctx context.Context, room domain.CreateRoomInput) (*domain.Room, error)
	Get(ctx context.Context, roomID string) (*domain.RoomInfo, error)
}
