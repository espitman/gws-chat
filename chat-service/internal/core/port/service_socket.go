package port

import "github.com/lxzan/gws"

/**
 * socketService implemented by service.socketService interface
 */

type SocketService interface {
	Subscribe(socket *gws.Conn)
}
