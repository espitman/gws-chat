package port

import (
	"context"

	"github.com/lxzan/gws"
)

/**
 * roomService implemented by service.roomService interface
 */

type RoomService interface {
	Create(socket *gws.Conn) error
	Subscribe(socket *gws.Conn) error
	GetSubscribers(roomID string) []*gws.Conn
	GetAudience(ctx context.Context, roomID string, userID uint32) (uint32, error)
}
