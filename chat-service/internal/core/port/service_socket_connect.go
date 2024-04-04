package port

import (
	"github.com/lxzan/gws"
	"net/http"
)

/**
 * messageService implemented by service.messageService interface
 */

type SocketConnetService interface {
	Open(writer http.ResponseWriter, request *http.Request, auth string, id string) (*gws.Conn, error)
}
