// This file generated automatically by gRPC gateway generator
// Generated at: 2024-04-05 23:06:57

package main

import (
	"github.com/gofiber/fiber/v2"
)

type router struct {
	userServiceHandler *userServiceHandler
}

func newRouter(
	userServiceHandler *userServiceHandler,
) *router {
	return &router{
		userServiceHandler: userServiceHandler,
	}
}

func (r *router) serve(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	r.swaggerRouter(app)
	r.userServiceRouter(v1.Group("/user-service"))
}
