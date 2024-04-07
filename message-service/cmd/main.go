package main

import (
	superConf "github.com/espitman/go-super-conf"
	"github.com/espitman/gws-chat/message-service/cmd/gRPC"
	"github.com/espitman/gws-chat/message-service/internal/adapter/database/postgres/pg"
	grpcclientuser "github.com/espitman/gws-chat/message-service/internal/adapter/grpcclient/user"
	"github.com/espitman/gws-chat/message-service/internal/core/service"
	validatorutil "github.com/espitman/gws-chat/pkg/util/validator"
)

func main() {
	validate := validatorutil.NewValidator()
	dsn := superConf.Get("db.postgres.dsn")
	pgDB := pg.NewDB(dsn)

	userRepositoryGrpc := grpcclientuser.NewGrpcClientUser()
	userService := service.NewUserService(userRepositoryGrpc)

	roomRepositoryPg := pg.NewRoomRepository(pgDB.Client)
	roomService := service.NewRoomService(
		roomRepositoryPg,
		userService,
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
