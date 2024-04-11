package grpc

import (
	"context"
	"strconv"
	"time"

	"github.com/espitman/gws-chat/message-service/internal/core/domain"
	pb "github.com/espitman/gws-chat/pkg/protos/protogen/message-service"
	timeutil "github.com/espitman/gws-chat/pkg/util/time"
)

func (h Handler) V1CreateRoom(ctx context.Context, req *pb.V1CreateRoomRequest) (*pb.V1CreateRoomResponse, error) {
	userIDCtx := ctx.Value("userID").(string)
	userID, _ := strconv.Atoi(userIDCtx)
	room := domain.CreateRoomInput{
		User1: uint32(userID),
		User2: uint32(req.UserId),
	}
	result, err := h.roomService.Crete(ctx, room)
	if err != nil {
		return nil, err
	}

	resp := pb.V1CreateRoomResponse{
		Id: result.RoomID,
	}
	return &resp, nil
}

func (h Handler) V1AddMemberToRoom(ctx context.Context, req *pb.V1AddMemberToRoomRequest) (*pb.V1AddMemberToRoomResponse, error) {
	userID := ctx.Value("userID").(string)
	user1, _ := strconv.Atoi(userID)
	user2 := req.UserId
	roomID := req.RoomID
	result, err := h.memberService.AddMember(ctx, roomID, uint32(user1), uint32(user2))
	if err != nil {
		return nil, err
	}
	var members []*pb.RoomMember
	for _, r := range result {
		members = append(members, &pb.RoomMember{
			Id:     int32(r.ID),
			RoomID: r.RoomID,
			UserID: int32(r.UserID),
		})
	}
	resp := pb.V1AddMemberToRoomResponse{
		Members: members,
	}
	return &resp, nil
}

func (h Handler) V1GetRoom(ctx context.Context, req *pb.V1GetRoomRequest) (*pb.V1GetRoomResponse, error) {
	//userIDCtx := ctx.Value("userID").(string)
	//userID, _ := strconv.Atoi(userIDCtx)

	roomInfo, err := h.roomService.Get(ctx, req.RoomID)
	if err != nil {
		return nil, err
	}

	var messages []*pb.Message
	messagesResult, err := h.messageService.GetRoomMessages(ctx, roomInfo)
	if err != nil {
		return nil, err
	}
	for _, m := range messagesResult {
		messages = append(messages, &pb.Message{
			Id:         m.ID,
			RoomID:     m.RoomID,
			UserID:     int32(m.UserID),
			UserName:   m.UserName,
			UserAvatar: m.UserAvatar,
			Body:       m.Body,
			Time:       timeutil.DateToString(m.Time),
			Type:       m.Type,
		})
	}

	return &pb.V1GetRoomResponse{
		RoomID: roomInfo.RoomID,
		User: &pb.MessageUser{
			Id:     int32(roomInfo.User.ID),
			Name:   roomInfo.User.Name,
			Avatar: roomInfo.User.Avatar,
			Status: roomInfo.User.Status,
		},
		Audience: &pb.MessageUser{
			Id:     int32(roomInfo.Audience.ID),
			Name:   roomInfo.Audience.Name,
			Avatar: roomInfo.Audience.Avatar,
			Status: roomInfo.Audience.Status,
		},
		Messages: messages,
	}, nil
}

func (h Handler) V1AddMessage(ctx context.Context, req *pb.V1AddMessageRequest) (*pb.V1AddMessageResponse, error) {
	userIDCtx := ctx.Value("userID").(string)
	userID, _ := strconv.Atoi(userIDCtx)

	msg := domain.Message{
		RoomID: req.RoomID,
		UserID: uint32(userID),
		Body:   req.Body,
		Time:   time.Now(),
	}
	result, err := h.messageService.Crete(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &pb.V1AddMessageResponse{
		Message: &pb.Message{
			Id:     result.ID,
			RoomID: result.RoomID,
			UserID: int32(result.UserID),
			Body:   result.Body,
			Time:   timeutil.DateToString(result.Time),
		},
	}, nil
}

func (h Handler) V1GetRoomMessages(ctx context.Context, req *pb.V1GetRoomMessagesRequest) (*pb.V1GetRoomMessagesResponse, error) {
	//userIDCtx := ctx.Value("userID").(string)
	//userID, _ := strconv.Atoi(userIDCtx)
	var messages []*pb.Message

	roomInfo, err := h.roomService.Get(ctx, req.RoomID)
	if err != nil {
		return nil, err
	}

	result, err := h.messageService.GetRoomMessages(ctx, roomInfo)
	if err != nil {
		return nil, err
	}
	for _, m := range result {
		messages = append(messages, &pb.Message{
			Id:     m.ID,
			RoomID: m.RoomID,
			UserID: int32(m.UserID),
			Body:   m.Body,
			Time:   timeutil.DateToString(m.Time),
		})
	}
	return &pb.V1GetRoomMessagesResponse{Messages: messages}, nil
}

func (h Handler) V1GetAudienceID(ctx context.Context, req *pb.V1GetAudienceIDRequest) (*pb.V1GetAudienceIDResponse, error) {
	userIDCtx := ctx.Value("userID").(string)
	userID, _ := strconv.Atoi(userIDCtx)

	audienceID, err := h.roomService.GetAudienceID(ctx, req.RoomID, uint32(userID))
	if err != nil {
		return nil, err
	}

	return &pb.V1GetAudienceIDResponse{
		RoomID:     req.RoomID,
		AudienceID: *audienceID,
	}, nil
}
