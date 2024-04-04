package memdb

import (
	"fmt"
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
		fmt.Println(err)
		return err
	}
	txn.Commit()
	return nil
}

func (s RoomRepository) Subscribe(roomID string, socket *gws.Conn) error {
	txn := s.subscriptionDb.Txn(true)
	subscriptionSchema := SubscriptionSchema{
		ID:     uuid.NewString(),
		RoomID: roomID,
		Socket: socket,
	}
	err := txn.Insert("subscriptions", subscriptionSchema)
	if err != nil {
		fmt.Println(err)
		return err
	}
	txn.Commit()
	return nil
}

func (s RoomRepository) GetSubscribers(roomID string) []*gws.Conn {
	var sockets []*gws.Conn
	txn := s.subscriptionDb.Txn(false)
	raws, err := txn.Get("subscriptions", "roomID", roomID)
	fmt.Println(raws)
	if err != nil {
		panic(err)
	}
	for obj := raws.Next(); obj != nil; obj = raws.Next() {
		s := obj.(SubscriptionSchema)
		sockets = append(sockets, s.Socket)
	}
	return sockets
}

//func (s RoomRepository) Get(roomID string) *gws.Conn {
//	txn := s.db.Txn(false)
//	raw, err := txn.First("rooms", "id", roomID)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("raw:::", raw)
//	return raw.(RoomSchema).Room
//}
//
//func (s RoomRepository) Delete(roomID string) {
//	txn := s.db.Txn(true)
//	err := txn.Delete("rooms", &RoomSchema{ID: roomID})
//	if err != nil {
//		log.Fatal(err)
//	}
//	txn.Commit()
//}
