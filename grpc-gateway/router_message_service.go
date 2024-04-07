// This file generated automatically by gRPC gateway generator
// Generated at: 2024-04-07 23:32:20

package main

import "github.com/gofiber/fiber/v2"

func (r *router) messageServiceRouter(v fiber.Router) {
	v.Post(
		"/v1-add-member-to-room",
		AuthMiddleware(r.userServiceClient),
		r.messageServiceHandler.V1AddMemberToRoom,
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
}
