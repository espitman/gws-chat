package pg

import (
	"github.com/espitman/gws-chat/user-service/internal/adapter/database/postgres/ent"
	"github.com/espitman/gws-chat/user-service/internal/core/domain"
)

func userSchemaToUserDomainPointerMapper(c *ent.User) *domain.User {
	return &domain.User{
		ID:       c.ID,
		Name:     c.Name,
		Password: c.Password,
		Avatar:   *c.Avatar,
		Status:   c.Status,
	}
}

func userSchemasToUserDomainsPointerMapper(cs []*ent.User) []*domain.User {
	resp := make([]*domain.User, len(cs))
	for i, c := range cs {
		resp[i] = userSchemaToUserDomainPointerMapper(c)
	}
	return resp
}

func userDomainToUserSchema(d domain.User) *ent.User {
	return &ent.User{
		Name: d.Name,
	}
}

func userDomainsToUserSchemasMapper(ds []domain.User) []*ent.User {
	resp := make([]*ent.User, len(ds))
	for i, d := range ds {
		resp[i] = userDomainToUserSchema(d)
	}
	return resp
}
