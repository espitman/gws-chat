// This file generated automatically by gRPC gateway generator
// Generated at: 2024-04-14 21:17:44

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
	v.Post(
		"/v1-get-audience-i-d",
		AuthMiddleware(r.userServiceClient),
		r.messageServiceHandler.V1GetAudienceID,
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
	v.Get(
		"/chats",
		AuthMiddleware(r.userServiceClient),
		r.messageServiceHandler.V1GetUserChats,
	)
}
