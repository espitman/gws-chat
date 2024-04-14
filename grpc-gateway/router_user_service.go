// This file generated automatically by gRPC gateway generator
// Generated at: 2024-04-14 21:17:44

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
		"/v1-get-by-i-ds",
		AuthMiddleware(r.userServiceClient),
		r.userServiceHandler.V1GetByIDs,
	)
	v.Post(
		"/login",
		r.userServiceHandler.V1Login,
	)
	v.Post(
		"/v1-set-online",
		AuthMiddleware(r.userServiceClient),
		r.userServiceHandler.V1SetOnline,
	)
}
