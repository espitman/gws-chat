package service

import (
	"context"

	"github.com/espitman/gws-chat/message-service/internal/core/domain"
	"github.com/espitman/gws-chat/message-service/internal/core/port"
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

func (s *UserService) Get(ctx context.Context, userID uint32) (*domain.User, error) {
	return s.userRepositoryGrpc.Get(ctx, userID)
}
