package port

import (
	"context"

	"github.com/espitman/gws-chat/message-service/internal/core/domain"
)

/**
 * memberRepository implemented by repository.memberRepository interface
 */

type MemberRepositoryPg interface {
	AddMember(ctx context.Context, roomID string, userID uint32) (*domain.Member, error)
}
