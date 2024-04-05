// This file generated automatically by gRPC gateway generator
// Generated at: 2024-04-06 02:01:18

package main

import (
	userpb "github.com/espitman/gws-chat/pkg/protos/protogen/user-service"
	"github.com/gofiber/fiber/v2"
)

type userServiceHandler struct {
	pb userpb.UserServiceClient
}

func newUserServiceHandler(pb userpb.UserServiceClient) *userServiceHandler {
	return &userServiceHandler{
		pb: pb,
	}
}

// V1GetAll
// @Summary V1GetAll
// @Description V1GetAll
// @Tags user_service
// @Produce json
// @Security BearerAuth
// @Param body body userpb.V1GetAllRequest true "body"
// @Success 200 {object} userServiceV1GetAllResponseDto
// @Router /api/v1/user-service/all [Get]
func (h *userServiceHandler) V1GetAll(c *fiber.Ctx) error {
	var fctx = fiberCtx{c}
	ctx := getCtx(fctx)
	var reqDto userpb.V1GetAllRequest
	_ = fctx.QueryParser(&reqDto)
	_ = fctx.ParamsParser(&reqDto)
	_ = fctx.BodyParser(&reqDto)
	res, err := h.pb.V1GetAll(ctx, &reqDto)
	if err != nil {
		return fctx.BadRequest(err)
	}
	return fctx.ResponseOk(res)
}

// V1Login
// @Summary V1Login
// @Description V1Login
// @Tags user_service
// @Produce json
// @Security BearerAuth
// @Param body body userpb.V1LoginRequest true "body"
// @Success 200 {object} userServiceV1LoginResponseDto
// @Router /api/v1/user-service/login [Post]
func (h *userServiceHandler) V1Login(c *fiber.Ctx) error {
	var fctx = fiberCtx{c}
	ctx := getCtx(fctx)
	var reqDto userpb.V1LoginRequest
	_ = fctx.QueryParser(&reqDto)
	_ = fctx.ParamsParser(&reqDto)
	_ = fctx.BodyParser(&reqDto)
	res, err := h.pb.V1Login(ctx, &reqDto)
	if err != nil {
		return fctx.BadRequest(err)
	}
	return fctx.ResponseOk(res)
}
