package pg

import (
	"context"
	"github.com/espitman/gws-chat/chat-service/internal/adapter/database/postgres/ent"
	"github.com/espitman/gws-chat/chat-service/internal/core/domain"
)

type MessageRepository struct {
	client *ent.Client
}

/**
 * messageRepository implements port.messageRepository interface
 */

func NewMessageRepository(client *ent.Client) *MessageRepository {
	return &MessageRepository{
		client: client,
	}
}

func (r *MessageRepository) Crete(ctx context.Context, d domain.Message) (*domain.Message, error) {
	newD, err := r.client.Message.
		Create().
		SetName(d.Name).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return messageSchemaToMessageDomainPointerMapper(newD), nil
}

func (r *MessageRepository) Get(ctx context.Context, ID int) (*domain.Message, error) {
	u, err := r.client.Message.Get(ctx, ID)
	if err != nil {
		return nil, err
	}
	return messageSchemaToMessageDomainPointerMapper(u), nil
}

func (r *MessageRepository) Update(ctx context.Context, ID int, d domain.Message) (*domain.Message, error) {
	u, err := r.client.Message.UpdateOneID(ID).SetName(d.Name).Save(ctx)
	if err != nil {
		return nil, err
	}
	return messageSchemaToMessageDomainPointerMapper(u), nil
}

func (r *MessageRepository) Delete(ctx context.Context, ID int) (*domain.Message, error) {
	u, err := r.client.Message.Get(ctx, ID)
	if err != nil {
		return nil, err
	}
	err = r.client.Message.DeleteOneID(ID).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return messageSchemaToMessageDomainPointerMapper(u), nil
}
