package memdb

import (
	"github.com/hashicorp/go-memdb"
	"github.com/lxzan/gws"
)

type RoomSchema struct {
	ID string
}

type SubscriptionSchema struct {
	ID     string
	RoomID string
	Socket *gws.Conn
}

func getRoomSchema() *memdb.DBSchema {
	return &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"rooms": {
				Name: "rooms",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
				},
			},
		},
	}
}

func getSubscriptionSchema() *memdb.DBSchema {
	return &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"subscriptions": {
				Name: "subscriptions",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"roomID": {
						Name:    "roomID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "RoomID"},
					},
				},
			},
		},
	}
}
