package pg

import (
	"context"

	"github.com/espitman/gws-chat/message-service/internal/adapter/database/postgres/ent"
	"github.com/espitman/gws-chat/message-service/internal/adapter/database/postgres/ent/room"
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

func (r *RoomRepository) Crete(ctx context.Context, d domain.Room) (*domain.Room, bool, error) {
	crRoom, err := r.client.Room.Query().Where(room.UsersEQ(d.Users)).First(ctx)
	if crRoom != nil {
		return roomSchemaToRoomDomainPointerMapper(crRoom), false, nil
	}
	newRoom, err := r.client.Room.
		Create().
		SetRoomID(d.RoomID).
		SetUsers(d.Users).
		Save(ctx)
	if err != nil {
		return nil, false, err
	}
	return roomSchemaToRoomDomainPointerMapper(newRoom), true, nil
}

func (r *RoomRepository) Get(ctx context.Context, roomID string) (*domain.Room, error) {
	aRoom, err := r.client.Room.Query().Where(room.RoomID(roomID)).First(ctx)
	if err != nil {
		return nil, err
	}
	return roomSchemaToRoomDomainPointerMapper(aRoom), nil
}

func (r *RoomRepository) GetRooms(ctx context.Context, IDs []string) ([]*domain.Room, error) {
	rooms, err := r.client.Room.Query().Where(room.RoomIDIn(IDs...)).All(ctx)
	if err != nil {
		return nil, err
	}
	return roomSchemasToRoomDomainsPointerMapper(rooms), nil
}
