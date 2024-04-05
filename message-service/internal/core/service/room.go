package service

import (
	"context"
	"github.com/espitman/gws-chat/message-service/internal/core/domain"
	"github.com/espitman/gws-chat/message-service/internal/core/port"
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

func (s *RoomService) Crete(ctx context.Context, room domain.Room) (*domain.Room, error) {
	var result domain.Room
	pgRoom, err := s.roomRepositoryPg.Crete(ctx, room)
	if err != nil {
		return nil, err
	}
	result.ID = pgRoom.ID
	result.Name = pgRoom.Name
	room.ID = pgRoom.ID
	return &result, nil
}

func (s *RoomService) Get(ctx context.Context, ID int) (*domain.Room, error) {
	var result domain.Room
	pgRoom, err := s.roomRepositoryPg.Get(ctx, ID)
	if err != nil {
		return nil, err
	}
	result.ID = pgRoom.ID
	result.Name = pgRoom.Name
	return &result, nil
}

func (s *RoomService) Update(ctx context.Context, ID int, room domain.Room) (*domain.Room, error) {
	var result domain.Room
	pgRoom, err := s.roomRepositoryPg.Update(ctx, ID, room)
	if err != nil {
		return nil, err
	}
	result = *pgRoom
	return &result, nil
}

func (s *RoomService) Delete(ctx context.Context, ID int) (*domain.Room, error) {
	var result domain.Room
	var err error
	pgRoom, err := s.roomRepositoryPg.Delete(ctx, ID)
	if err != nil {
		return nil, err
	}
	result = *pgRoom
	return &result, nil
}
