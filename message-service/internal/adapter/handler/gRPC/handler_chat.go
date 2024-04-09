package grpc

import (
	"context"
	"strconv"

	pb "github.com/espitman/gws-chat/pkg/protos/protogen/message-service"
	timeutil "github.com/espitman/gws-chat/pkg/util/time"
)

func (h Handler) V1GetUserChats(ctx context.Context, _ *pb.V1GetUserChatsRequest) (*pb.V1GetUserChatsResponse, error) {
	userIDCtx := ctx.Value("userID").(string)
	userID, _ := strconv.Atoi(userIDCtx)
	chats, err := h.chatService.GetUserChats(ctx, uint32(userID))
	if err != nil {
		return nil, err
	}
	var resp pb.V1GetUserChatsResponse
	for _, chat := range chats {
		resp.Chats = append(resp.Chats, &pb.UserChat{
			User: &pb.MessageUser{
				Id:     int32(chat.User.ID),
				Name:   chat.User.Name,
				Avatar: chat.User.Avatar,
				Status: chat.User.Status,
			},
			Message: &pb.UserChatMessage{
				RoomID: chat.RoomID,
				Body:   chat.Body,
				Time:   timeutil.DateToString(chat.Time),
			},
		})
	}
	return &resp, nil
}
