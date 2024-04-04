package main

import (
	superConf "github.com/espitman/go-super-conf"
	"github.com/espitman/gws-chat/chat-service/cmd/api"
	"github.com/espitman/gws-chat/chat-service/internal/adapter/database/memdb"
	"github.com/espitman/gws-chat/chat-service/internal/adapter/database/postgres/pg"
	"github.com/espitman/gws-chat/chat-service/internal/core/service"
	validatorutil "github.com/espitman/gws-chat/pkg/util/validator"
)

func main() {
	validate := validatorutil.NewValidator()
	dsn := superConf.Get("db.postgres.dsn")
	pgDB := pg.NewDB(dsn)
	messageRepositoryPg := pg.NewMessageRepository(pgDB.Client)
	messageService := service.NewMessageService(
		messageRepositoryPg,
	)

	socketRepositoryMD := memdb.NewSocketRepository()
	socketService := service.NewSocketService(socketRepositoryMD)

	roomRepositoryMD := memdb.NewRoomRepository()
	roomService := service.NewRoomService(roomRepositoryMD)

	socketConnectService := service.NewSocketConnectService(socketService, roomService)

	// +salvation NewRepository
	api.Run(
		validate,
		messageService,
		socketConnectService,
		socketService,
		roomService,
		// +salvation RunServiceAPI
	)
}
