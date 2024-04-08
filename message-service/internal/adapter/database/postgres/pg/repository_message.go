package pg

import (
	"context"

	"github.com/espitman/gws-chat/message-service/internal/adapter/database/postgres/ent"
	"github.com/espitman/gws-chat/message-service/internal/adapter/database/postgres/ent/message"
	"github.com/espitman/gws-chat/message-service/internal/core/domain"
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
		SetRoomID(d.RoomID).
		SetUserID(d.UserID).
		SetBody(d.Body).
		SetTime(d.Time).
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

func (r *MessageRepository) GetRoomMessages(ctx context.Context, roomID string) ([]*domain.Message, error) {
	messages, err := r.client.Message.Query().Where(message.RoomID(roomID)).All(ctx)
	if err != nil {
		return nil, err
	}
	return messageSchemasToMessageDomainsPointerMapper(messages), nil
}
