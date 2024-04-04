package service

import (
	"github.com/espitman/gws-chat/chat-service/internal/adapter/handler/socket"
	"github.com/espitman/gws-chat/chat-service/internal/core/port"
	"github.com/google/uuid"
	"github.com/lxzan/gws"
	"net/http"
)

/**
 * socketService implements port.socketService interface
 */

type SocketConnectService struct {
	upgrader      *gws.Upgrader
	socketService port.SocketService
	roomService   port.RoomService
}

func NewSocketConnectService(socketService port.SocketService, roomService port.RoomService) *SocketConnectService {
	sh := socket.NewHandler(socketService, roomService)
	upgrader := gws.NewUpgrader(&sh, &gws.ServerOption{
		ParallelEnabled:   true,                                 // Parallel message processing
		Recovery:          gws.Recovery,                         // Exception recovery
		PermessageDeflate: gws.PermessageDeflate{Enabled: true}, // Enable compression
	})
	return &SocketConnectService{
		upgrader:      upgrader,
		socketService: socketService,
	}
}

func (s *SocketConnectService) Open(writer http.ResponseWriter, request *http.Request, userID string, roomID string) (*gws.Conn, error) {
	socketID := uuid.NewString()
	iSocket, err := s.upgrader.Upgrade(writer, request)
	if err != nil {
		return nil, err
	}
	iSocket.Session().Store("socketID", socketID)
	iSocket.Session().Store("roomID", roomID)
	return iSocket, nil
}
