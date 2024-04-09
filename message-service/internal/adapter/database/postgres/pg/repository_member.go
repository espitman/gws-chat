package pg

import (
	"context"

	"github.com/espitman/gws-chat/message-service/internal/adapter/database/postgres/ent"
	"github.com/espitman/gws-chat/message-service/internal/adapter/database/postgres/ent/member"
	"github.com/espitman/gws-chat/message-service/internal/core/domain"
)

type MemberRepository struct {
	client *ent.Client
}

/**
 * memberRepository implements port.memberRepository interface
 */

func NewMemberRepository(client *ent.Client) *MemberRepository {
	return &MemberRepository{
		client: client,
	}
}

func (r *MemberRepository) AddMember(ctx context.Context, roomID string, userID uint32) (*domain.Member, error) {
	newD, err := r.client.Member.
		Create().
		SetRoomID(roomID).
		SetUserID(userID).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return memberSchemaToMemberDomainPointerMapper(newD), nil
}

func (r *MemberRepository) GetUserAllRooms(ctx context.Context, userID uint32) ([]*domain.Member, error) {
	result, err := r.client.Member.Query().Where(member.UserID(userID)).All(ctx)
	if err != nil {
		return nil, err
	}
	return memberSchemasToMemberDomainsPointerMapper(result), nil
}
