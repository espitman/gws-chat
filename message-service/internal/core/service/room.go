package service

import (
	"context"
	"errors"
	"sort"
	"strconv"
	"strings"

	"github.com/espitman/gws-chat/message-service/internal/core/domain"
	"github.com/espitman/gws-chat/message-service/internal/core/port"
	"github.com/google/uuid"
)

/**
 * RoomService implements port.RoomService interface
 */

type RoomService struct {
	roomRepositoryPg port.RoomRepositoryPg
	userService      port.UserService
}

func NewRoomService(
	roomRepositoryPg port.RoomRepositoryPg,
	userService port.UserService,
) *RoomService {
	return &RoomService{
		roomRepositoryPg: roomRepositoryPg,
		userService:      userService,
	}
}

func (s *RoomService) generateRoomUsers(user1 uint32, user2 uint32) string {
	ids := []uint32{user1, user2}
	sort.Slice(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})
	result := strconv.Itoa(int(ids[0])) + "," + strconv.Itoa(int(ids[1])) + ","
	return result
}

func (s *RoomService) isRoomMember(userID uint32, users string) bool {
	allUsers := strings.Split(users, ",")
	for _, user := range allUsers {
		id, _ := strconv.Atoi(user)
		if uint32(id) == userID {
			return true
		}
	}
	return false
}

func (s *RoomService) getUseIDFromCtx(ctx context.Context) uint32 {
	userIDString := ctx.Value("userID").(string)
	userID, _ := strconv.Atoi(userIDString)
	return uint32(userID)
}

func (s *RoomService) getAudienceID(userID uint32, users string) uint32 {
	allUsers := strings.Split(users, ",")
	for _, user := range allUsers {
		id, _ := strconv.Atoi(user)
		if uint32(id) != userID {
			return uint32(id)
		}
	}
	return 0
}

func (s *RoomService) Crete(ctx context.Context, roomInput domain.CreateRoomInput) (*domain.Room, error) {
	var result domain.Room
	room := domain.Room{
		RoomID: uuid.NewString(),
		Users:  s.generateRoomUsers(roomInput.User1, roomInput.User2),
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

func (s *RoomService) Get(ctx context.Context, roomID string) (*domain.RoomInfo, error) {
	userID := s.getUseIDFromCtx(ctx)
	pgRoom, err := s.roomRepositoryPg.Get(ctx, roomID)
	if err != nil {
		return nil, err
	}

	isMember := s.isRoomMember(userID, pgRoom.Users)
	if !isMember {
		return nil, errors.New("forbidden")
	}

	user, err := s.userService.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	audienceID := s.getAudienceID(userID, pgRoom.Users)
	audience, err := s.userService.Get(ctx, audienceID)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return &domain.RoomInfo{
		ID:     pgRoom.ID,
		RoomID: pgRoom.RoomID,
		User: domain.User{
			ID:     user.ID,
			Name:   user.Name,
			Avatar: user.Avatar,
			Status: user.Status,
		},
		Audience: domain.User{
			ID:     audience.ID,
			Name:   audience.Name,
			Avatar: audience.Avatar,
			Status: audience.Status,
		},
	}, nil
}
