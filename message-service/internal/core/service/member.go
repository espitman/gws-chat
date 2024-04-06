package service

import (
	"context"

	"github.com/espitman/gws-chat/message-service/internal/core/domain"
	"github.com/espitman/gws-chat/message-service/internal/core/port"
)

/**
 * MemberService implements port.MemberService interface
 */

type MemberService struct {
	memberRepositoryPg port.MemberRepositoryPg
}

func NewMemberService(
	memberRepositoryPg port.MemberRepositoryPg,
) *MemberService {
	return &MemberService{
		memberRepositoryPg: memberRepositoryPg,
	}
}

func (s *MemberService) AddMember(ctx context.Context, roomID string, user1 uint32, user2 uint32) ([]*domain.Member, error) {
	var result []*domain.Member
	pgMember1, err := s.memberRepositoryPg.AddMember(ctx, roomID, user1)
	if err != nil {
		return nil, err
	}
	result = append(result, &domain.Member{
		ID:     pgMember1.ID,
		RoomID: pgMember1.RoomID,
		UserID: pgMember1.UserID,
	})
	pgMember2, err := s.memberRepositoryPg.AddMember(ctx, roomID, user2)
	if err != nil {
		return nil, err
	}
	result = append(result, &domain.Member{
		ID:     pgMember2.ID,
		RoomID: pgMember2.RoomID,
		UserID: pgMember2.UserID,
	})
	return result, nil
}
