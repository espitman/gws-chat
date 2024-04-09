package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/espitman/gws-chat/user-service/internal/core/domain"
	"github.com/espitman/gws-chat/user-service/internal/core/port"
	"github.com/golang-jwt/jwt/v5"
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
	user.Password = s.passwordGenerator(user.Password)
	user.Avatar = s.avatar(user.Name)
	return s.userRepositoryPg.Crete(ctx, user)
}

func (s *UserService) avatar(name string) string {
	return "https://api.dicebear.com/8.x/avataaars/svg?seed=" + name
}

func generateJWTToken(user domain.User) (string, error) {
	hmacSampleSecret := []byte("jabamaChat")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Name,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func parseJWTToken(tokenString string) (*domain.User, error) {
	var resp domain.User
	hmacSampleSecret := []byte("jabamaChat")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		resp.ID = int(claims["id"].(float64))
		resp.Name = claims["username"].(string)
	} else {
		return nil, err
	}
	return &resp, nil
}

func (s *UserService) Login(ctx context.Context, user domain.User) (*domain.User, error) {
	pgUser, err := s.userRepositoryPg.GetByName(ctx, user)
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
	resp := domain.User{
		ID:     pgUser.ID,
		Name:   pgUser.Name,
		Avatar: pgUser.Avatar,
		Status: pgUser.Status,
	}
	token, _ := generateJWTToken(resp)
	resp.Token = token
	return &resp, nil
}

func (s *UserService) GetAll(ctx context.Context, me bool, userID uint32) ([]*domain.User, error) {
	return s.userRepositoryPg.GetAll(ctx, me, int(userID))
}

func (s *UserService) ValidateToken(ctx context.Context, token string) (*domain.User, error) {
	user, err := parseJWTToken(token)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) Get(ctx context.Context, userID uint32) (*domain.User, error) {
	pgUser, err := s.userRepositoryPg.Get(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &domain.User{
		ID:     pgUser.ID,
		Name:   pgUser.Name,
		Avatar: pgUser.Avatar,
		Status: pgUser.Status,
	}, nil
}

func (s *UserService) GetByIDs(ctx context.Context, userIDs []uint32) ([]*domain.User, error) {
	var ids []int
	for _, d := range userIDs {
		ids = append(ids, int(d))
	}
	return s.userRepositoryPg.GetByIDs(ctx, ids)
}
