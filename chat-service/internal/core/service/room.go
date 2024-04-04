package service

import (
	"fmt"
	"github.com/lxzan/gws"
)

type RoomService struct {
}

func NewRoomService() *RoomService {
	return &RoomService{}
}

func (s *RoomService) Subscribe(socket *gws.Conn) {
	roomID, _ := socket.Session().Load("roomID")
	fmt.Println(">>>roomID:", roomID)
}
