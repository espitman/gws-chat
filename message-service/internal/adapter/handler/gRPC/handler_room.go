package grpc

import (
	"context"
	pb "github.com/espitman/gws-chat/pkg/protos/protogen/message-service"
)

func (h Handler) V1CreateRoom(ctx context.Context, req *pb.V1CreateRoomRequest) (*pb.V1CreateRoomResponse, error) {
	var resp pb.V1CreateRoomResponse
	return &resp, nil
}
