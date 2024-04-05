package main

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func (r *router) swaggerRouter(v fiber.Router) {
	v.Get("/swagger/*", swagger.New(swagger.Config{
		URL: "/swagger/doc.json",
	}))

	v.Get("/swagger/doc.json", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(swagger.New())
	})
}
