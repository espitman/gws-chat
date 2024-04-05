// This file generated automatically by gRPC gateway generator
// Generated at: 2024-04-06 00:28:32

package main

import (
	userpb "github.com/espitman/gws-chat/pkg/protos/protogen/user-service"
	"github.com/gofiber/fiber/v2"
)

type router struct {
	userServiceClient  userpb.UserServiceClient
	userServiceHandler *userServiceHandler
}

func newRouter(
	userServiceClient userpb.UserServiceClient,
	userServiceHandler *userServiceHandler,
) *router {
	return &router{
		userServiceClient:  userServiceClient,
		userServiceHandler: userServiceHandler,
	}
}

func (r *router) serve(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	r.swaggerRouter(app)
	r.userServiceRouter(v1.Group("/user-service"))
}
