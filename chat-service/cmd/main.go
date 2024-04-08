package main

import (
	"github.com/espitman/gws-chat/chat-service/cmd/api"
	"github.com/espitman/gws-chat/chat-service/internal/adapter/database/memdb"
	grpcclientmessage "github.com/espitman/gws-chat/chat-service/internal/adapter/grpcclient/message"
	grpcclientuser "github.com/espitman/gws-chat/chat-service/internal/adapter/grpcclient/user"
	"github.com/espitman/gws-chat/chat-service/internal/core/service"
	validatorutil "github.com/espitman/gws-chat/pkg/util/validator"
)

func main() {
	validate := validatorutil.NewValidator()

	userRepositoryGrpc := grpcclientuser.NewGrpcClientUser()
	userService := service.NewUserService(userRepositoryGrpc)

	messageRepositoryGrpc := grpcclientmessage.NewGrpcClientMessage()
	messageService := service.NewMessageService(*messageRepositoryGrpc)

	socketRepositoryMD := memdb.NewSocketRepository()
	socketService := service.NewSocketService(socketRepositoryMD)

	roomRepositoryMD := memdb.NewRoomRepository()
	roomService := service.NewRoomService(roomRepositoryMD)

	socketConnectService := service.NewSocketConnectService(socketService, roomService, messageService)

	// +salvation NewRepository
	api.Run(
		validate,
		messageService,
		socketConnectService,
		socketService,
		roomService,
		userService,
		// +salvation RunServiceAPI
	)
}
