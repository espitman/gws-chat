package port

import "github.com/lxzan/gws"

/**
 * socketService implemented by service.socketService interface
 */

type SocketService interface {
	Save(socket *gws.Conn)
	Get(socketID string) *gws.Conn
	Delete(socketID string)
}
