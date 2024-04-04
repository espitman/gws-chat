package fiberutil

import (
	"github.com/gofiber/fiber/v2"
)

type FiberCtx struct {
	*fiber.Ctx
}

type FiberOk struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Payload any    `json:"payload"`
}

type FiberError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (c *FiberCtx) error(code int, err error) error {
	fe := FiberError{
		Code:    code,
		Message: err.Error(),
	}
	return c.Status(code).JSON(fe)
}

func (c *FiberCtx) BadRequest(err error) error {
	return c.error(fiber.StatusBadRequest, err)
}

func (c *FiberCtx) Conflict(err error) error {
	return c.error(fiber.StatusConflict, err)
}

func (c *FiberCtx) ResponseOk(payload any) error {
	fo := FiberOk{
		Code:    fiber.StatusOK,
		Message: "ok",
		Payload: payload,
	}
	return c.JSON(fo)
}
