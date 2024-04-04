package port

import "github.com/lxzan/gws"

/**
 * roomService implemented by service.roomService interface
 */

type RoomService interface {
	Subscribe(room *gws.Conn)
}
