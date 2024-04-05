// This file generated automatically by gRPC gateway generator
// Generated at: 2024-04-05 19:02:50

package main

import (
	superConf "github.com/espitman/go-super-conf"
	_ "github.com/espitman/gws-chat/grpc-gateway/docs"
	"github.com/gofiber/fiber/v2"
	"log"
)

// @title           gRPC Gateway
// @version         1.0
// @description     This is gRPC Gateway
// @contact.name   API Support
// @contact.email  s.heidar@jabama.com
// @BasePath  /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	app := fiber.New()
	userServiceClient := newUserServiceClient()
	userServiceHandler := newUserServiceHandler(userServiceClient)
	router := newRouter(
		userServiceHandler,
	)
	router.serve(app)
	PORT := superConf.Get("app.port")
	log.Fatal(app.Listen(":" + PORT))
}
