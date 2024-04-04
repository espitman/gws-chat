package service

import (
	"fmt"
	"github.com/lxzan/gws"
)

/**
 * socketService implements port.socketService interface
 */

type SocketService struct {
}

func NewSocketService() *SocketService {
	return &SocketService{}
}

func (s *SocketService) Subscribe(socket *gws.Conn) {
	socketID, _ := socket.Session().Load("socketID")
	fmt.Println(">>>socketID:", socketID)
}
