package memdb

import (
	"github.com/hashicorp/go-memdb"
	"github.com/lxzan/gws"
)

type SocketSchema struct {
	ID     string
	Socket *gws.Conn
}

func getSocketSchema() *memdb.DBSchema {
	return &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"sockets": {
				Name: "sockets",
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
