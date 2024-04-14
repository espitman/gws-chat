// This file generated automatically by gRPC gateway generator
// Generated at: 2024-04-14 21:17:45

package main

import (
	superConf "github.com/espitman/go-super-conf"
	_ "github.com/espitman/gws-chat/grpc-gateway/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	app.Use(cors.New())
	userServiceClient := newUserServiceClient()
	userServiceHandler := newUserServiceHandler(userServiceClient)
	messageServiceClient := newMessageServiceClient()
	messageServiceHandler := newMessageServiceHandler(messageServiceClient)
	router := newRouter(
		userServiceClient,
		userServiceHandler,
		messageServiceHandler,
	)
	router.serve(app)
	PORT := superConf.Get("app.port")
	log.Fatal(app.Listen(":" + PORT))
}
