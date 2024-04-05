package pg

import (
	"context"
	"github.com/espitman/gws-chat/message-service/internal/adapter/database/postgres/ent"
	"github.com/espitman/gws-chat/message-service/internal/core/domain"
)

type RoomRepository struct {
	client *ent.Client
}

/**
 * roomRepository implements port.roomRepository interface
 */

func NewRoomRepository(client *ent.Client) *RoomRepository {
	return &RoomRepository{
		client: client,
	}
}

func (r *RoomRepository) Crete(ctx context.Context, d domain.Room) (*domain.Room, error) {
	newD, err := r.client.Room.
		Create().
		SetName(d.Name).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return roomSchemaToRoomDomainPointerMapper(newD), nil
}

func (r *RoomRepository) Get(ctx context.Context, ID int) (*domain.Room, error) {
	u, err := r.client.Room.Get(ctx, ID)
	if err != nil {
		return nil, err
	}
	return roomSchemaToRoomDomainPointerMapper(u), nil
}

func (r *RoomRepository) Update(ctx context.Context, ID int, d domain.Room) (*domain.Room, error) {
	u, err := r.client.Room.UpdateOneID(ID).SetName(d.Name).Save(ctx)
	if err != nil {
		return nil, err
	}
	return roomSchemaToRoomDomainPointerMapper(u), nil
}

func (r *RoomRepository) Delete(ctx context.Context, ID int) (*domain.Room, error) {
	u, err := r.client.Room.Get(ctx, ID)
	if err != nil {
		return nil, err
	}
	err = r.client.Room.DeleteOneID(ID).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return roomSchemaToRoomDomainPointerMapper(u), nil
}
