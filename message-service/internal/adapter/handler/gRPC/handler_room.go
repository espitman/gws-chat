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
	userID := ctx.Value("userID").(string)
	user1, _ := strconv.Atoi(userID)
	room := domain.CreateRoomInput{
		User1: uint32(user1),
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
	result, err := h.roomService.Get(ctx, req.RoomID)
	if err != nil {
		return nil, err
	}
	return &pb.V1GetRoomResponse{
		RoomID:     result.RoomID,
		UserID:     int32(result.UserID),
		UserName:   result.UserName,
		UserAvatar: result.UserAvatar,
		UserStatus: result.UserStatus,
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
		Id:     result.ID,
		RoomID: result.RoomID,
		UserId: int32(result.UserID),
		Body:   result.Body,
		Time:   timeutil.DateToString(result.Time),
	}, nil
}
