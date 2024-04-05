package grpc

import (
	"context"
	"github.com/espitman/gws-chat/message-service/internal/core/domain"
	pb "github.com/espitman/gws-chat/pkg/protos/protogen/message-service"
	"strconv"
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
