package main

import (
	superConf "github.com/espitman/go-super-conf"
	"github.com/espitman/gws-chat/message-service/cmd/gRPC"
	"github.com/espitman/gws-chat/message-service/internal/adapter/database/postgres/pg"
	"github.com/espitman/gws-chat/message-service/internal/core/service"
	validatorutil "github.com/espitman/gws-chat/pkg/util/validator"
)

func main() {
	validate := validatorutil.NewValidator()
	dsn := superConf.Get("db.postgres.dsn")
	pgDB := pg.NewDB(dsn)
	roomRepositoryPg := pg.NewRoomRepository(pgDB.Client)
	roomService := service.NewRoomService(
		roomRepositoryPg,
	)

	memberRepositoryPg := pg.NewMemberRepository(pgDB.Client)
	memberService := service.NewMemberService(
		memberRepositoryPg,
	)

	// +salvation NewRepository
	gRPC.Run(
		validate,
		roomService,
		memberService,
		// +salvation RunServiceGRPC
	)
}
