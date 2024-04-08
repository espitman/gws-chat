package service

import (
	"context"

	"github.com/espitman/gws-chat/chat-service/internal/core/domain"
	"github.com/espitman/gws-chat/chat-service/internal/core/port"
)

/**
 * UserService implements port.UserService interface
 */

type UserService struct {
	userRepositoryGrpc port.UserRepositoryGrpc
}

func NewUserService(
	userRepositoryGrpc port.UserRepositoryGrpc,
) *UserService {
	return &UserService{
		userRepositoryGrpc: userRepositoryGrpc,
	}
}

func (s *UserService) Validate(ctx context.Context, token string) (*domain.User, error) {
	return s.userRepositoryGrpc.Validate(ctx, token)
}
