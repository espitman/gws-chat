// This file generated automatically by gRPC gateway generator
// Generated at: 2024-04-06 02:01:18

package main

import "github.com/gofiber/fiber/v2"

func (r *router) messageServiceRouter(v fiber.Router) {
	v.Post(
		"/room",
		AuthMiddleware(r.userServiceClient),
		r.messageServiceHandler.V1CreateRoom,
	)
}
