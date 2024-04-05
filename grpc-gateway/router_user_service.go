// This file generated automatically by gRPC gateway generator
// Generated at: 2024-04-06 00:28:30

package main

import "github.com/gofiber/fiber/v2"

func (r *router) userServiceRouter(v fiber.Router) {
	v.Get(
		"/all",
		AuthMiddleware(r.userServiceClient),
		r.userServiceHandler.V1GetAll,
	)
	v.Post(
		"/login",
		r.userServiceHandler.V1Login,
	)
	v.Post(
		"/v1-validate-token",
		AuthMiddleware(r.userServiceClient),
		r.userServiceHandler.V1ValidateToken,
	)
}
