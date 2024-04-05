package main

import (
	superConf "github.com/espitman/go-super-conf"
	validatorutil "github.com/espitman/gws-chat/pkg/util/validator"
	"github.com/espitman/gws-chat/user-service/cmd/gRPC"
	"github.com/espitman/gws-chat/user-service/internal/adapter/database/postgres/pg"
	"github.com/espitman/gws-chat/user-service/internal/core/service"
)

func main() {
	validate := validatorutil.NewValidator()
	dsn := superConf.Get("db.postgres.dsn")
	pgDB := pg.NewDB(dsn)
	userRepositoryPg := pg.NewUserRepository(pgDB.Client)
	userService := service.NewUserService(
		userRepositoryPg,
	)

	// +salvation NewRepository
	gRPC.Run(
		validate,
		userService,
		// +salvation RunServiceGRPC
	)
}
