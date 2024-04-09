// This file generated automatically by gRPC gateway generator
// Generated at: 2024-04-10 01:10:28

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

// V1Get
// @Summary V1Get
// @Description V1Get
// @Tags user_service
// @Produce json
// @Security BearerAuth
// @Param userID path string true "userID"
// @Success 200 {object} userServiceV1GetResponseDto
// @Router /api/v1/user-service/{userID} [Get]
func (h *userServiceHandler) V1Get(c *fiber.Ctx) error {
	var fctx = fiberCtx{c}
	ctx := getCtx(fctx)
	var reqDto userpb.V1GetRequest
	_ = fctx.ParamsParser(&reqDto)
	res, err := h.pb.V1Get(ctx, &reqDto)
	if err != nil {
		return fctx.BadRequest(err)
	}
	return fctx.ResponseOk(res)
}

// V1GetAll
// @Summary V1GetAll
// @Description V1GetAll
// @Tags user_service
// @Produce json
// @Security BearerAuth
// @Param me query string false "me"
// @Success 200 {object} userServiceV1GetAllResponseDto
// @Router /api/v1/user-service/all [Get]
func (h *userServiceHandler) V1GetAll(c *fiber.Ctx) error {
	var fctx = fiberCtx{c}
	ctx := getCtx(fctx)
	var reqDto userpb.V1GetAllRequest
	_ = fctx.QueryParser(&reqDto)
	res, err := h.pb.V1GetAll(ctx, &reqDto)
	if err != nil {
		return fctx.BadRequest(err)
	}
	return fctx.ResponseOk(res)
}

// V1GetByIDs
// @Summary V1GetByIDs
// @Description V1GetByIDs
// @Tags user_service
// @Produce json
// @Security BearerAuth
// @Param body body userpb.V1GetByIDsRequest true "body"
// @Success 200 {object} userServiceV1GetByIDsResponseDto
// @Router /api/v1/user-service/v1-get-by-i-ds [Post]
func (h *userServiceHandler) V1GetByIDs(c *fiber.Ctx) error {
	var fctx = fiberCtx{c}
	ctx := getCtx(fctx)
	var reqDto userpb.V1GetByIDsRequest
	_ = fctx.QueryParser(&reqDto)
	_ = fctx.ParamsParser(&reqDto)
	_ = fctx.BodyParser(&reqDto)
	res, err := h.pb.V1GetByIDs(ctx, &reqDto)
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
