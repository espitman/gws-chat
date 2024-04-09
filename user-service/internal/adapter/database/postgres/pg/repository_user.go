package pg

import (
	"context"

	"github.com/espitman/gws-chat/user-service/internal/adapter/database/postgres/ent"
	"github.com/espitman/gws-chat/user-service/internal/adapter/database/postgres/ent/user"
	"github.com/espitman/gws-chat/user-service/internal/core/domain"
)

type UserRepository struct {
	client *ent.Client
}

/**
 * userRepository implements port.userRepository interface
 */

func NewUserRepository(client *ent.Client) *UserRepository {
	return &UserRepository{
		client: client,
	}
}

func (r *UserRepository) Crete(ctx context.Context, d domain.User) (*domain.User, error) {
	newD, err := r.client.User.
		Create().
		SetName(d.Name).
		SetPassword(d.Password).
		SetAvatar(d.Avatar).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return userSchemaToUserDomainPointerMapper(newD), nil
}

func (r *UserRepository) Get(ctx context.Context, userID uint32) (*domain.User, error) {
	u, err := r.client.User.Get(ctx, int(userID))
	if err != nil {
		return nil, err
	}
	return userSchemaToUserDomainPointerMapper(u), nil
}

func (r *UserRepository) GetByName(ctx context.Context, d domain.User) (*domain.User, error) {
	u, err := r.client.User.Query().Where(user.Name(d.Name)).First(ctx)
	if err != nil {
		return nil, err
	}
	return userSchemaToUserDomainPointerMapper(u), nil
}

func (r *UserRepository) GetAll(ctx context.Context, me bool, userID int) ([]*domain.User, error) {
	q := r.client.User.Query()
	if !me {
		q.Where(user.IDNEQ(userID))
	}
	users, err := q.All(ctx)
	if err != nil {
		return nil, err
	}
	return userSchemasToUserDomainsPointerMapper(users), nil
}

//
//func (r *UserRepository) Update(ctx context.Context, ID int, d domain.User) (*domain.User, error) {
//	u, err := r.client.User.UpdateOneID(ID).SetName(d.Name).Save(ctx)
//	if err != nil {
//		return nil, err
//	}
//	return userSchemaToUserDomainPointerMapper(u), nil
//}
//
//func (r *UserRepository) Delete(ctx context.Context, ID int) (*domain.User, error) {
//	u, err := r.client.User.GetByName(ctx, ID)
//	if err != nil {
//		return nil, err
//	}
//	err = r.client.User.DeleteOneID(ID).Exec(ctx)
//	if err != nil {
//		return nil, err
//	}
//	return userSchemaToUserDomainPointerMapper(u), nil
//}
