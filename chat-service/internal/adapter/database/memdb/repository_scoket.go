package memdb

import (
	"fmt"
	"github.com/hashicorp/go-memdb"
	"github.com/lxzan/gws"
	"log"
)

type SocketRepository struct {
	db *memdb.MemDB
}

func NewSocketRepository() *SocketRepository {
	schema := getSocketSchema()
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}
	return &SocketRepository{
		db: db,
	}
}

func (s SocketRepository) Save(socket *gws.Conn) error {
	txn := s.db.Txn(true)
	socketID, _ := socket.Session().Load("socketID")
	socketSchema := SocketSchema{
		ID:     socketID.(string),
		Socket: socket,
	}
	err := txn.Insert("sockets", socketSchema)
	if err != nil {
		fmt.Println(err)
		return err
	}
	txn.Commit()

	return nil
}

func (s SocketRepository) Get(socketID string) *gws.Conn {
	txn := s.db.Txn(false)
	raw, err := txn.First("sockets", "id", socketID)
	if err != nil {
		panic(err)
	}
	fmt.Println("raw:::", raw)
	return raw.(SocketSchema).Socket
}

func (s SocketRepository) Delete(socketID string) {
	txn := s.db.Txn(true)
	err := txn.Delete("sockets", &SocketSchema{ID: socketID})
	if err != nil {
		log.Fatal(err)
	}
	txn.Commit()
}
