package socket

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/espitman/gws-chat/chat-service/internal/core/domain"
	"github.com/espitman/gws-chat/chat-service/internal/core/port"
	"github.com/lxzan/gws"
)

const (
	PingInterval = 50 * time.Second
	PingWait     = 100 * time.Second
)

type Handler struct {
	pubSub         *gochannel.GoChannel
	socketService  port.SocketService
	roomService    port.RoomService
	messageService port.MessageService
}

func NewHandler(pubSub *gochannel.GoChannel, socketService port.SocketService, roomService port.RoomService, messageService port.MessageService) Handler {
	return Handler{
		pubSub:         pubSub,
		socketService:  socketService,
		roomService:    roomService,
		messageService: messageService,
	}
}

func (h *Handler) OnOpen(socket *gws.Conn) {
	_ = socket.SetDeadline(time.Now().Add(PingInterval + PingWait))
	socketID, _ := socket.Session().Load("socketID")
	socket.WriteString(socketID.(string))
	fmt.Println("OnOpen")
}

func (h *Handler) OnClose(socket *gws.Conn, err error) {
	roomID, _ := socket.Session().Load("roomID")
	socketID, _ := socket.Session().Load("socketID")
	h.socketService.Delete(socketID.(string))
	h.roomService.Delete(roomID.(string))
	h.roomService.UnSubscribe(roomID.(string))
	fmt.Println("OnClose", roomID, socketID)
}

func (h *Handler) OnPing(socket *gws.Conn, payload []byte) {
	_ = socket.SetDeadline(time.Now().Add(PingInterval + PingWait))
	_ = socket.WritePong(nil)
	fmt.Println("OnPing")
}

func (h *Handler) OnPong(socket *gws.Conn, payload []byte) {
	fmt.Println("OnPong")
}

func (h *Handler) OnMessage(socket *gws.Conn, message *gws.Message) {
	defer message.Close()
	ctx := context.TODO()
	roomID, _ := socket.Session().Load("roomID")
	socketID, _ := socket.Session().Load("socketID")
	userID, _ := socket.Session().Load("userID")
	//token, _ := socket.Session().Load("token")
	fmt.Println("OnMessage", roomID, socketID, message.Data, userID)

	var messageBody domain.MessageBody
	err := json.Unmarshal(message.Data.Bytes(), &messageBody)
	if err != nil {
		fmt.Println(err)
	}

	u, _ := strconv.Atoi(userID.(string))
	h.messageService.Text(ctx, message, socketID.(string), roomID.(string), uint32(u), messageBody)
}
