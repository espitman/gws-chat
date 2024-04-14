package memdb

import (
	"log"

	"github.com/espitman/gws-chat/chat-service/internal/core/domain"
	"github.com/google/uuid"
	"github.com/hashicorp/go-memdb"
	"github.com/lxzan/gws"
)

type RoomRepository struct {
	roomDb         *memdb.MemDB
	subscriptionDb *memdb.MemDB
}

func NewRoomRepository() *RoomRepository {
	roomSchema := getRoomSchema()
	subscriptionSchema := getSubscriptionSchema()
	roomDb, err := memdb.NewMemDB(roomSchema)
	if err != nil {
		panic(err)
	}
	subscriptionDb, err := memdb.NewMemDB(subscriptionSchema)
	if err != nil {
		panic(err)
	}
	return &RoomRepository{
		roomDb:         roomDb,
		subscriptionDb: subscriptionDb,
	}
}

func (s RoomRepository) Create(roomID string) error {
	txn := s.roomDb.Txn(true)
	roomSchema := RoomSchema{
		ID: roomID,
	}
	err := txn.Insert("rooms", roomSchema)
	if err != nil {
		log.Println(err)
		return err
	}
	txn.Commit()
	return nil
}

func (s RoomRepository) Subscribe(roomID string, socket *gws.Conn) error {
	txn := s.subscriptionDb.Txn(true)
	id := uuid.NewString()
	subscriptionSchema := SubscriptionSchema{
		ID:     id,
		RoomID: roomID,
		Socket: socket,
	}
	err := txn.Insert("subscriptions", subscriptionSchema)
	if err != nil {
		log.Println(err)
		return err
	}
	txn.Commit()
	return nil
}

func (s RoomRepository) GetSubscribers(roomID string) []*gws.Conn {
	var sockets []*gws.Conn
	txn := s.subscriptionDb.Txn(false)
	raws, err := txn.Get("subscriptions", "roomID", roomID)
	if err != nil {
		log.Println(err)
	}
	for obj := raws.Next(); obj != nil; obj = raws.Next() {
		s := obj.(SubscriptionSchema)
		sockets = append(sockets, s.Socket)
	}
	return sockets
}

func (s RoomRepository) Delete(roomID string) {
	txn := s.roomDb.Txn(true)
	err := txn.Delete("rooms", &RoomSchema{ID: roomID})
	if err != nil {
		log.Println(err)
	}
	txn.Commit()
}

func (s RoomRepository) UnSubscribe(roomID string) {
	txn := s.subscriptionDb.Txn(true)
	_, err := txn.DeleteAll("subscriptions", "roomID", roomID)
	if err != nil {
		log.Println(err)
	}
	txn.Commit()
}

func (s RoomRepository) GetAllRooms() []domain.Room {
	var rooms []domain.Room
	txn := s.roomDb.Txn(false)
	raws, err := txn.Get("rooms", "id")
	if err != nil {
		log.Println(err)
	}
	for obj := raws.Next(); obj != nil; obj = raws.Next() {
		s := obj.(RoomSchema)
		r := domain.Room{
			ID: s.ID,
		}
		rooms = append(rooms, r)
	}
	return rooms
}

func (s RoomRepository) GetAllSubscribers() []domain.Subscriber {
	var subscribers []domain.Subscriber
	txn := s.subscriptionDb.Txn(false)
	raws, err := txn.Get("subscriptions", "id")
	if err != nil {
		log.Println(err)
	}
	for obj := raws.Next(); obj != nil; obj = raws.Next() {
		s := obj.(SubscriptionSchema)
		r := domain.Subscriber{
			ID:     s.ID,
			RoomID: s.RoomID,
		}
		subscribers = append(subscribers, r)
	}
	return subscribers
}
