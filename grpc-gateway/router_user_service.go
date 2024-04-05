// This file generated automatically by gRPC gateway generator
// Generated at: 2024-04-05 16:15:24

package main

import "github.com/gofiber/fiber/v2"

func (r *router) userServiceRouter(v fiber.Router) {
	v.Post(
		"/v1-login",
		r.userServiceHandler.V1Login,
	)
}
