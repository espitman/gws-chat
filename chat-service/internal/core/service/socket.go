package service

import (
	"fmt"
	"github.com/espitman/gws-chat/chat-service/internal/core/port"
	"github.com/lxzan/gws"
)

/**
 * socketService implements port.socketService interface
 */

type SocketService struct {
	socketRepositoryMD port.SocketRepositoryMD
}

func NewSocketService(socketRepositoryMD port.SocketRepositoryMD) *SocketService {
	return &SocketService{
		socketRepositoryMD: socketRepositoryMD,
	}
}

func (s *SocketService) Save(socket *gws.Conn) {
	socketID, _ := socket.Session().Load("socketID")
	err := s.socketRepositoryMD.Save(socket)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(">>>socketID:", socketID)
}

func (s *SocketService) Get(socketID string) *gws.Conn {
	return s.socketRepositoryMD.Get(socketID)
}

func (s *SocketService) Delete(socketID string) {
	s.socketRepositoryMD.Delete(socketID)
}
