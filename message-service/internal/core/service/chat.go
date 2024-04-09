package service

import (
	"context"
	"strconv"
	"strings"

	"github.com/espitman/gws-chat/message-service/internal/core/domain"
	"github.com/espitman/gws-chat/message-service/internal/core/port"
)

/**
 * ChatService implements port.ChatService interface
 */

type ChatService struct {
	memberRepositoryPg  port.MemberRepositoryPg
	messageRepositoryPg port.MessageRepositoryPg
	userRepositoryGrpc  port.UserRepositoryGrpc
	roomRepositoryPg    port.RoomRepositoryPg
}

func NewChatService(
	memberRepositoryPg port.MemberRepositoryPg,
	messageRepositoryPg port.MessageRepositoryPg,
	userRepositoryGrpc port.UserRepositoryGrpc,
	roomRepositoryPg port.RoomRepositoryPg,
) *ChatService {
	return &ChatService{
		memberRepositoryPg:  memberRepositoryPg,
		messageRepositoryPg: messageRepositoryPg,
		userRepositoryGrpc:  userRepositoryGrpc,
		roomRepositoryPg:    roomRepositoryPg,
	}
}

func (s *ChatService) getUserChatsRoomsIDs(ctx context.Context, userID uint32) ([]string, error) {
	rooms, err := s.memberRepositoryPg.GetUserAllRooms(ctx, userID)
	if err != nil {
		return nil, err
	}
	var roomIDs []string
	for _, room := range rooms {
		roomIDs = append(roomIDs, room.RoomID)
	}
	return roomIDs, err
}

func (s *ChatService) getAudienceID(userID uint32, users string) uint32 {
	allUsers := strings.Split(users, ",")
	for _, user := range allUsers {
		id, _ := strconv.Atoi(user)
		if uint32(id) != userID {
			return uint32(id)
		}
	}
	return 0
}

func (s *ChatService) getRoomsUser(ctx context.Context, userID uint32, roomIDs []string) (map[string]*domain.User, error) {
	rooms, err := s.roomRepositoryPg.GetRooms(ctx, roomIDs)
	if err != nil {
		return nil, err
	}
	var userIDs []uint32
	roomsUser := make(map[string]uint32)
	for _, room := range rooms {
		audienceID := s.getAudienceID(userID, room.Users)
		userIDs = append(userIDs, audienceID)
		roomsUser[room.RoomID] = audienceID
	}
	users, err := s.userRepositoryGrpc.GetByIds(ctx, userIDs)
	if err != nil {
		return nil, err
	}
	usersData := make(map[uint32]*domain.User)
	for _, user := range users {
		usersData[user.ID] = user
	}

	result := make(map[string]*domain.User)
	for roomID, userID := range roomsUser {
		result[roomID] = usersData[userID]
	}
	return result, nil
}

func (s *ChatService) GetUserChats(ctx context.Context, userID uint32) ([]*domain.Chat, error) {
	roomIDs, err := s.getUserChatsRoomsIDs(ctx, userID)
	if err != nil {
		return nil, err
	}

	messages, err := s.messageRepositoryPg.GetRoomLastMessages(ctx, userID, roomIDs...)
	if err != nil {
		return nil, err
	}

	roomsUser, err := s.getRoomsUser(ctx, userID, roomIDs)
	if err != nil {
		return nil, err
	}

	var chats []*domain.Chat
	for _, message := range messages {
		chats = append(chats, &domain.Chat{
			User: domain.User{
				ID:     message.UserID,
				Name:   roomsUser[message.RoomID].Name,
				Avatar: roomsUser[message.RoomID].Avatar,
				Status: roomsUser[message.RoomID].Status,
			},
			RoomID: message.RoomID,
			Body:   message.Body,
			Time:   message.Time,
		})
	}
	return chats, nil
}
