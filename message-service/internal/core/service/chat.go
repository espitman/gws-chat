package service

import (
	"context"

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
}

func NewChatService(
	memberRepositoryPg port.MemberRepositoryPg,
	messageRepositoryPg port.MessageRepositoryPg,
	userRepositoryGrpc port.UserRepositoryGrpc,
) *ChatService {
	return &ChatService{
		memberRepositoryPg:  memberRepositoryPg,
		messageRepositoryPg: messageRepositoryPg,
		userRepositoryGrpc:  userRepositoryGrpc,
	}
}

func (s *ChatService) GetUserChats(ctx context.Context, userID uint32) ([]*domain.Chat, error) {
	rooms, err := s.memberRepositoryPg.GetUserAllRooms(ctx, userID)
	if err != nil {
		return nil, err
	}
	var roomIDs []string
	for _, room := range rooms {
		roomIDs = append(roomIDs, room.RoomID)
	}
	var chats []*domain.Chat
	messages, err := s.messageRepositoryPg.GetRoomLastMessages(ctx, userID, roomIDs...)

	var userIDs []uint32
	for _, message := range messages {
		userIDs = append(userIDs, message.UserID)
	}
	users, err := s.userRepositoryGrpc.GetByIds(ctx, userIDs)
	if err != nil {
		return nil, err
	}

	usersData := make(map[uint32]*domain.User)
	for _, user := range users {
		usersData[user.ID] = user
	}

	for _, message := range messages {
		chats = append(chats, &domain.Chat{
			User: domain.User{
				ID:     message.UserID,
				Name:   usersData[userID].Name,
				Avatar: usersData[userID].Avatar,
				Status: usersData[userID].Status,
			},
			RoomID: message.RoomID,
			Body:   message.Body,
			Time:   message.Time,
		})
	}
	return chats, nil
}
