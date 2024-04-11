package port

import (
	"net/http"

	"github.com/lxzan/gws"
)

/**
 * messageService implemented by service.messageService interface
 */

type SocketConnetService interface {
	Open(writer http.ResponseWriter, request *http.Request, userID string, roomID string, token string) (*gws.Conn, error)
}
