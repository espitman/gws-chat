// This file generated automatically by gRPC gateway generator
// Generated at: 2024-04-14 21:17:44

package main

import (
	messagepb "github.com/espitman/gws-chat/pkg/protos/protogen/message-service"
	"github.com/gofiber/fiber/v2"
)

type messageServiceHandler struct {
	pb messagepb.MessageServiceClient
}

func newMessageServiceHandler(pb messagepb.MessageServiceClient) *messageServiceHandler {
	return &messageServiceHandler{
		pb: pb,
	}
}

// V1AddMemberToRoom
// @Summary V1AddMemberToRoom
// @Description V1AddMemberToRoom
// @Tags message_service
// @Produce json
// @Security BearerAuth
// @Param body body messagepb.V1AddMemberToRoomRequest true "body"
// @Success 200 {object} messageServiceV1AddMemberToRoomResponseDto
// @Router /api/v1/message-service/v1-add-member-to-room [Post]
func (h *messageServiceHandler) V1AddMemberToRoom(c *fiber.Ctx) error {
	var fctx = fiberCtx{c}
	ctx := getCtx(fctx)
	var reqDto messagepb.V1AddMemberToRoomRequest
	_ = fctx.QueryParser(&reqDto)
	_ = fctx.ParamsParser(&reqDto)
	_ = fctx.BodyParser(&reqDto)
	res, err := h.pb.V1AddMemberToRoom(ctx, &reqDto)
	if err != nil {
		return fctx.BadRequest(err)
	}
	return fctx.ResponseOk(res)
}

// V1AddMessage
// @Summary V1AddMessage
// @Description V1AddMessage
// @Tags message_service
// @Produce json
// @Security BearerAuth
// @Param body body messagepb.V1AddMessageRequest true "body"
// @Success 200 {object} messageServiceV1AddMessageResponseDto
// @Router /api/v1/message-service/v1-add-message [Post]
func (h *messageServiceHandler) V1AddMessage(c *fiber.Ctx) error {
	var fctx = fiberCtx{c}
	ctx := getCtx(fctx)
	var reqDto messagepb.V1AddMessageRequest
	_ = fctx.QueryParser(&reqDto)
	_ = fctx.ParamsParser(&reqDto)
	_ = fctx.BodyParser(&reqDto)
	res, err := h.pb.V1AddMessage(ctx, &reqDto)
	if err != nil {
		return fctx.BadRequest(err)
	}
	return fctx.ResponseOk(res)
}

// V1CreateRoom
// @Summary V1CreateRoom
// @Description V1CreateRoom
// @Tags message_service
// @Produce json
// @Security BearerAuth
// @Param body body messagepb.V1CreateRoomRequest true "body"
// @Success 200 {object} messageServiceV1CreateRoomResponseDto
// @Router /api/v1/message-service/room [Post]
func (h *messageServiceHandler) V1CreateRoom(c *fiber.Ctx) error {
	var fctx = fiberCtx{c}
	ctx := getCtx(fctx)
	var reqDto messagepb.V1CreateRoomRequest
	_ = fctx.BodyParser(&reqDto)
	res, err := h.pb.V1CreateRoom(ctx, &reqDto)
	if err != nil {
		return fctx.BadRequest(err)
	}
	return fctx.ResponseOk(res)
}

// V1GetAudienceID
// @Summary V1GetAudienceID
// @Description V1GetAudienceID
// @Tags message_service
// @Produce json
// @Security BearerAuth
// @Param body body messagepb.V1GetAudienceIDRequest true "body"
// @Success 200 {object} messageServiceV1GetAudienceIDResponseDto
// @Router /api/v1/message-service/v1-get-audience-i-d [Post]
func (h *messageServiceHandler) V1GetAudienceID(c *fiber.Ctx) error {
	var fctx = fiberCtx{c}
	ctx := getCtx(fctx)
	var reqDto messagepb.V1GetAudienceIDRequest
	_ = fctx.QueryParser(&reqDto)
	_ = fctx.ParamsParser(&reqDto)
	_ = fctx.BodyParser(&reqDto)
	res, err := h.pb.V1GetAudienceID(ctx, &reqDto)
	if err != nil {
		return fctx.BadRequest(err)
	}
	return fctx.ResponseOk(res)
}

// V1GetRoom
// @Summary V1GetRoom
// @Description V1GetRoom
// @Tags message_service
// @Produce json
// @Security BearerAuth
// @Param RoomID path string true "RoomID"
// @Success 200 {object} messageServiceV1GetRoomResponseDto
// @Router /api/v1/message-service/room/{RoomID} [Get]
func (h *messageServiceHandler) V1GetRoom(c *fiber.Ctx) error {
	var fctx = fiberCtx{c}
	ctx := getCtx(fctx)
	var reqDto messagepb.V1GetRoomRequest
	_ = fctx.ParamsParser(&reqDto)
	res, err := h.pb.V1GetRoom(ctx, &reqDto)
	if err != nil {
		return fctx.BadRequest(err)
	}
	return fctx.ResponseOk(res)
}

// V1GetRoomMessages
// @Summary V1GetRoomMessages
// @Description V1GetRoomMessages
// @Tags message_service
// @Produce json
// @Security BearerAuth
// @Param RoomID path string true "RoomID"
// @Success 200 {object} messageServiceV1GetRoomMessagesResponseDto
// @Router /api/v1/message-service/room/{RoomID}/messages [Get]
func (h *messageServiceHandler) V1GetRoomMessages(c *fiber.Ctx) error {
	var fctx = fiberCtx{c}
	ctx := getCtx(fctx)
	var reqDto messagepb.V1GetRoomMessagesRequest
	_ = fctx.ParamsParser(&reqDto)
	res, err := h.pb.V1GetRoomMessages(ctx, &reqDto)
	if err != nil {
		return fctx.BadRequest(err)
	}
	return fctx.ResponseOk(res)
}

// V1GetUserChats
// @Summary V1GetUserChats
// @Description V1GetUserChats
// @Tags message_service
// @Produce json
// @Security BearerAuth
// @Success 200 {object} messageServiceV1GetUserChatsResponseDto
// @Router /api/v1/message-service/chats [Get]
func (h *messageServiceHandler) V1GetUserChats(c *fiber.Ctx) error {
	var fctx = fiberCtx{c}
	ctx := getCtx(fctx)
	var reqDto messagepb.V1GetUserChatsRequest
	res, err := h.pb.V1GetUserChats(ctx, &reqDto)
	if err != nil {
		return fctx.BadRequest(err)
	}
	return fctx.ResponseOk(res)
}
