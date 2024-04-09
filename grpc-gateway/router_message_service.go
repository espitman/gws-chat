// This file generated automatically by gRPC gateway generator
// Generated at: 2024-04-09 12:09:49

package main

import "github.com/gofiber/fiber/v2"

func (r *router) messageServiceRouter(v fiber.Router) {
	v.Post(
		"/v1-add-member-to-room",
		AuthMiddleware(r.userServiceClient),
		r.messageServiceHandler.V1AddMemberToRoom,
	)
	v.Post(
		"/v1-add-message",
		AuthMiddleware(r.userServiceClient),
		r.messageServiceHandler.V1AddMessage,
	)
	v.Post(
		"/room",
		AuthMiddleware(r.userServiceClient),
		r.messageServiceHandler.V1CreateRoom,
	)
	v.Get(
		"/room/:RoomID",
		AuthMiddleware(r.userServiceClient),
		r.messageServiceHandler.V1GetRoom,
	)
	v.Get(
		"/room/:RoomID/messages",
		AuthMiddleware(r.userServiceClient),
		r.messageServiceHandler.V1GetRoomMessages,
	)
}
