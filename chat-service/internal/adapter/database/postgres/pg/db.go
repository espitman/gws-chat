package pg

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/espitman/gws-chat/chat-service/internal/adapter/database/postgres/ent"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type DB struct {
	Client *ent.Client
}

type contextKey string

const (
	transactionKey contextKey = "tx"
)

func NewDB(dsn string) *DB {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv))

	if err := client.Schema.Create(context.Background()); err != nil {
		panic(err)
	}

	return &DB{
		Client: client,
	}

}
