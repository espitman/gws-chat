package port

import (
	"context"

	"github.com/espitman/gws-chat/message-service/internal/core/domain"
)

/**
 * MemberService implemented by service.MemberService interface
 */

type MemberService interface {
	AddMember(ctx context.Context, roomID string, user1 uint32, user2 uint32) ([]*domain.Member, error)
}
