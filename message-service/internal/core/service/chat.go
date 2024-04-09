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
}

func NewChatService(
	memberRepositoryPg port.MemberRepositoryPg,
	messageRepositoryPg port.MessageRepositoryPg,
) *ChatService {
	return &ChatService{
		memberRepositoryPg:  memberRepositoryPg,
		messageRepositoryPg: messageRepositoryPg,
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
	for _, message := range messages {
		chats = append(chats, &domain.Chat{
			User: domain.User{
				ID:     0,
				Name:   "",
				Avatar: "",
				Status: "",
			},
			RoomID: message.RoomID,
			Body:   message.Body,
			Time:   message.Time,
		})
	}

	return chats, nil
}
