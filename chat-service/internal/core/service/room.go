package service

import (
	"github.com/espitman/gws-chat/chat-service/internal/core/port"
	"github.com/lxzan/gws"
)

type RoomService struct {
	roomRepositoryMD port.RoomRepositoryMD
}

func NewRoomService(roomRepositoryMD port.RoomRepositoryMD) *RoomService {
	return &RoomService{
		roomRepositoryMD: roomRepositoryMD,
	}
}

func (s *RoomService) Create(socket *gws.Conn) error {
	roomID, _ := socket.Session().Load("roomID")
	return s.roomRepositoryMD.Create(roomID.(string))
}

func (s *RoomService) Subscribe(socket *gws.Conn) error {
	roomID, _ := socket.Session().Load("roomID")
	return s.roomRepositoryMD.Subscribe(roomID.(string), socket)
}

func (s *RoomService) GetSubscribers(roomID string) []*gws.Conn {
	return s.roomRepositoryMD.GetSubscribers(roomID)
}
