package pg

import (
	"github.com/espitman/gws-chat/chat-service/internal/adapter/database/postgres/ent"
	"github.com/espitman/gws-chat/chat-service/internal/core/domain"
)

func messageSchemaToMessageDomainPointerMapper(c *ent.Message) *domain.Message {
	return &domain.Message{
    ID:     c.ID,
    Name:   c.Name,
	}
}

func messageSchemasToMessageDomainsPointerMapper(cs []*ent.Message) []*domain.Message {
	resp := make([]*domain.Message, len(cs))
	for i, c := range cs {
		resp[i] = messageSchemaToMessageDomainPointerMapper(c)
	}
	return resp
}

func messageDomainToMessageSchema(d domain.Message) *ent.Message {
	return &ent.Message{
		Name:            d.Name,
	}
}

func messageDomainsToMessageSchemasMapper(ds []domain.Message) []*ent.Message {
	resp := make([]*ent.Message, len(ds))
	for i, d := range ds {
		resp[i] = messageDomainToMessageSchema(d)
	}
	return resp
}
