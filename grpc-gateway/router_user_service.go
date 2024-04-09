// This file generated automatically by gRPC gateway generator
// Generated at: 2024-04-09 12:09:49

package main

import "github.com/gofiber/fiber/v2"

func (r *router) userServiceRouter(v fiber.Router) {
	v.Get(
		"/all",
		AuthMiddleware(r.userServiceClient),
		r.userServiceHandler.V1GetAll,
	)
	v.Get(
		"/:userID",
		AuthMiddleware(r.userServiceClient),
		r.userServiceHandler.V1Get,
	)
	v.Post(
		"/login",
		r.userServiceHandler.V1Login,
	)
}
