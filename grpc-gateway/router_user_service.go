// This file generated automatically by gRPC gateway generator
// Generated at: 2024-04-05 19:02:48

package main

import "github.com/gofiber/fiber/v2"

func (r *router) userServiceRouter(v fiber.Router) {
	v.Get(
		"/all",
		AuthMiddleware,
		r.userServiceHandler.V1GetAll,
	)
	v.Post(
		"/login",
		r.userServiceHandler.V1Login,
	)
}
