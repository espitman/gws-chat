package port

import "github.com/lxzan/gws"

/**
 * SocketRepositoryMD implemented by memdb.SocketRepositoryMD interface
 */

type SocketRepositoryMD interface {
	Save(socket *gws.Conn) error
	Get(socketID string) *gws.Conn
	Delete(socketID string)
}
