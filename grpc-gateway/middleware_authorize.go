package main

import (
	"fmt"
	userpb "github.com/espitman/gws-chat/pkg/protos/protogen/user-service"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func AuthMiddleware(userServiceClient userpb.UserServiceClient) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")
		if authorization == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  fiber.StatusUnauthorized,
				"message": "Unauthorized",
			})
		}

		token, err := userServiceClient.V1ValidateToken(c.Context(), &userpb.V1ValidateTokenRequest{
			Token: strings.Replace(authorization, "Bearer ", "", 1),
		})
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  fiber.StatusUnauthorized,
				"message": "Unauthorized",
			})
		}

		fmt.Println(token)

		return c.Next()
	}
}
