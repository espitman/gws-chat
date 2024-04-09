// This file generated automatically by gRPC gateway generator
// Generated at: 2024-04-10 01:10:28

package main

import "github.com/gofiber/fiber/v2"

func (r *router) userServiceRouter(v fiber.Router) {
	v.Get(
		"/:userID",
		AuthMiddleware(r.userServiceClient),
		r.userServiceHandler.V1Get,
	)
	v.Get(
		"/all",
		AuthMiddleware(r.userServiceClient),
		r.userServiceHandler.V1GetAll,
	)
	v.Post(
		"/v1-get-by-i-ds",
		AuthMiddleware(r.userServiceClient),
		r.userServiceHandler.V1GetByIDs,
	)
	v.Post(
		"/login",
		r.userServiceHandler.V1Login,
	)
}
