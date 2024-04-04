package port

import "github.com/lxzan/gws"

/**
 * RoomRepositoryMD implemented by memdb.RoomRepositoryMD interface
 */

type RoomRepositoryMD interface {
	Create(roomID string) error
	Subscribe(roomID string, socket *gws.Conn) error
	GetSubscribers(roomID string) []*gws.Conn
	//Get(roomID string) *gws.Conn
	//Delete(roomID string)
}
