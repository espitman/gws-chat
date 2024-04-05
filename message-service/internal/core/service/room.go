package service

import (
	"context"
	"github.com/espitman/gws-chat/message-service/internal/core/domain"
	"github.com/espitman/gws-chat/message-service/internal/core/port"
	"github.com/google/uuid"
	"sort"
	"strconv"
)

/**
 * RoomService implements port.RoomService interface
 */

type RoomService struct {
	roomRepositoryPg port.RoomRepositoryPg
}

func NewRoomService(
	roomRepositoryPg port.RoomRepositoryPg,
) *RoomService {
	return &RoomService{
		roomRepositoryPg: roomRepositoryPg,
	}
}

func generateRoomUsers(user1 uint32, user2 uint32) string {
	ids := []uint32{user1, user2}
	sort.Slice(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})
	result := strconv.Itoa(int(ids[0])) + "," + strconv.Itoa(int(ids[1])) + ","
	return result
}

func (s *RoomService) Crete(ctx context.Context, roomInput domain.CreateRoomInput) (*domain.Room, error) {
	var result domain.Room
	room := domain.Room{
		RoomID: uuid.NewString(),
		Users:  generateRoomUsers(roomInput.User1, roomInput.User2),
	}
	pgRoom, err := s.roomRepositoryPg.Crete(ctx, room)
	if err != nil {
		return nil, err
	}
	result.ID = pgRoom.ID
	result.RoomID = pgRoom.RoomID
	result.Users = pgRoom.Users
	return &result, nil
}
