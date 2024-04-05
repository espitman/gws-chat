package pg

import (
	"github.com/espitman/gws-chat/message-service/internal/adapter/database/postgres/ent"
	"github.com/espitman/gws-chat/message-service/internal/core/domain"
)

func roomSchemaToRoomDomainPointerMapper(c *ent.Room) *domain.Room {
	return &domain.Room{
		ID:     c.ID,
		RoomID: c.RoomID,
		Users:  c.Users,
	}
}

func roomSchemasToRoomDomainsPointerMapper(cs []*ent.Room) []*domain.Room {
	resp := make([]*domain.Room, len(cs))
	for i, c := range cs {
		resp[i] = roomSchemaToRoomDomainPointerMapper(c)
	}
	return resp
}

func roomDomainToRoomSchema(d domain.Room) *ent.Room {
	return &ent.Room{
		ID:     d.ID,
		RoomID: d.RoomID,
		Users:  d.Users,
	}
}

func roomDomainsToRoomSchemasMapper(ds []domain.Room) []*ent.Room {
	resp := make([]*ent.Room, len(ds))
	for i, d := range ds {
		resp[i] = roomDomainToRoomSchema(d)
	}
	return resp
}
