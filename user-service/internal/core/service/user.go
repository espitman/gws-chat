package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/espitman/gws-chat/user-service/internal/core/domain"
	"github.com/espitman/gws-chat/user-service/internal/core/port"
	"golang.org/x/crypto/bcrypt"
	"time"
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
	return "https://api.dicebear.com/8.x/thumbs/svg?seed=" + name
}

func generateJWTToken(user domain.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["username"] = user.Name
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiration time
	secretKey := []byte("jabamaChat")
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func parseJWTToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("jabamaChat"), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
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
	token, _ := generateJWTToken(user)
	return &domain.User{
		ID:     pgUser.ID,
		Name:   pgUser.Name,
		Avatar: pgUser.Avatar,
		Status: pgUser.Status,
		Token:  token,
	}, nil
}

func (s *UserService) GetAll(ctx context.Context) ([]*domain.User, error) {
	return s.userRepositoryPg.GetAll(ctx)
}
