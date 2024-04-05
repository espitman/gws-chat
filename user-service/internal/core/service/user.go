package service

import (
	"context"
	"errors"
	"github.com/espitman/gws-chat/user-service/internal/core/domain"
	"github.com/espitman/gws-chat/user-service/internal/core/port"
	"golang.org/x/crypto/bcrypt"
)

/**
 * UserService implements port.UserService interface
 */

type UserService struct {
	userRepositoryPg port.UserRepositoryPg
}

func NewUserService(
	userRepositoryPg port.UserRepositoryPg,
) *UserService {
	return &UserService{
		userRepositoryPg: userRepositoryPg,
	}
}

func (s *UserService) passwordGenerator(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

func (s *UserService) passwordCompare(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (s *UserService) register(ctx context.Context, user domain.User) (*domain.User, error) {
	newUser := user
	newUser.Password = s.passwordGenerator(user.Password)
	return s.userRepositoryPg.Crete(ctx, newUser)
}

func (s *UserService) Login(ctx context.Context, user domain.User) (*domain.User, error) {
	pgUser, err := s.userRepositoryPg.Get(ctx, user)
	if err != nil {
		if err.Error() != "ent: user not found" {
			return nil, err
		}
		pgUser, err = s.register(ctx, user)
		if err != nil {
			return nil, err
		}
	}
	err = s.passwordCompare(pgUser.Password, user.Password)
	if err != nil {
		return nil, errors.New("invalid password")
	}
	return &domain.User{
		ID:     pgUser.ID,
		Name:   pgUser.Name,
		Avatar: pgUser.Avatar,
		Status: pgUser.Status,
	}, nil
}

func (s *UserService) GetAll(ctx context.Context) ([]*domain.User, error) {
	return s.userRepositoryPg.GetAll(ctx)
}
