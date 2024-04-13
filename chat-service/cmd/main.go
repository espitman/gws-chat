package main

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/espitman/gws-chat/chat-service/cmd/api"
	"github.com/espitman/gws-chat/chat-service/cmd/subsctib"
	"github.com/espitman/gws-chat/chat-service/internal/adapter/database/memdb"
	grpcclientmessage "github.com/espitman/gws-chat/chat-service/internal/adapter/grpcclient/message"
	grpcclientuser "github.com/espitman/gws-chat/chat-service/internal/adapter/grpcclient/user"
	"github.com/espitman/gws-chat/chat-service/internal/core/service"
	validatorutil "github.com/espitman/gws-chat/pkg/util/validator"
)

func main() {
	validate := validatorutil.NewValidator()

	pubSub := gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewStdLogger(false, false),
	)

	userRepositoryGrpc := grpcclientuser.NewGrpcClientUser()
	userService := service.NewUserService(userRepositoryGrpc)

	messageRepositoryGrpc := grpcclientmessage.NewGrpcClientMessage()
	messageService := service.NewMessageService(*messageRepositoryGrpc)

	socketRepositoryMD := memdb.NewSocketRepository()
	socketService := service.NewSocketService(socketRepositoryMD)

	roomRepositoryMD := memdb.NewRoomRepository()
	roomService := service.NewRoomService(roomRepositoryMD, messageRepositoryGrpc)

	socketConnectService := service.NewSocketConnectService(pubSub, socketService, roomService, messageService)

	pusherService := service.NewPusherService()

	go subsctib.Run(pubSub, pusherService)

	// +salvation NewRepository
	api.Run(
		validate,
		pubSub,
		messageService,
		socketConnectService,
		socketService,
		roomService,
		userService,
		// +salvation RunServiceAPI
	)
}
