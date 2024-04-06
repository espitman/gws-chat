package pg

import (
	"github.com/espitman/gws-chat/message-service/internal/adapter/database/postgres/ent"
	"github.com/espitman/gws-chat/message-service/internal/core/domain"
)

func memberSchemaToMemberDomainPointerMapper(c *ent.Member) *domain.Member {
	return &domain.Member{
		ID:     c.ID,
		RoomID: c.RoomID,
		UserID: c.UserID,
	}
}

func memberSchemasToMemberDomainsPointerMapper(cs []*ent.Member) []*domain.Member {
	resp := make([]*domain.Member, len(cs))
	for i, c := range cs {
		resp[i] = memberSchemaToMemberDomainPointerMapper(c)
	}
	return resp
}

func memberDomainToMemberSchema(d domain.Member) *ent.Member {
	return &ent.Member{
		ID:     d.ID,
		RoomID: d.RoomID,
		UserID: d.UserID,
	}
}

func memberDomainsToMemberSchemasMapper(ds []domain.Member) []*ent.Member {
	resp := make([]*ent.Member, len(ds))
	for i, d := range ds {
		resp[i] = memberDomainToMemberSchema(d)
	}
	return resp
}
