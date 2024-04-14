package service

import (
	"context"

	"github.com/espitman/gws-chat/chat-service/internal/core/domain"
	"github.com/espitman/gws-chat/chat-service/internal/core/port"
	"github.com/lxzan/gws"
)

type RoomService struct {
	roomRepositoryMD      port.RoomRepositoryMD
	messageRepositoryGrpc port.MessageRepositoryGrpc
}

func NewRoomService(roomRepositoryMD port.RoomRepositoryMD, messageRepositoryGrpc port.MessageRepositoryGrpc) *RoomService {
	return &RoomService{
		roomRepositoryMD:      roomRepositoryMD,
		messageRepositoryGrpc: messageRepositoryGrpc,
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

func (s *RoomService) GetAudience(ctx context.Context, roomID string, userID uint32) (uint32, error) {
	return s.messageRepositoryGrpc.GetAudienceID(ctx, roomID, userID)
}

func (s *RoomService) Delete(roomID string) {
	s.roomRepositoryMD.Delete(roomID)
}

func (s *RoomService) UnSubscribe(roomID string) {
	s.roomRepositoryMD.UnSubscribe(roomID)
}

func (s *RoomService) GetAllRooms() []domain.Room {
	return s.roomRepositoryMD.GetAllRooms()
}

func (s *RoomService) GetAllSubscribers() []domain.Subscriber {
	return s.roomRepositoryMD.GetAllSubscribers()
}
