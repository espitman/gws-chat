package socket

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/espitman/gws-chat/chat-service/internal/core/domain"
	"github.com/espitman/gws-chat/chat-service/internal/core/port"
	"github.com/lxzan/gws"
)

const (
	PingInterval = 50 * time.Second
	PingWait     = 100 * time.Second
)

type Handler struct {
	socketService  port.SocketService
	roomService    port.RoomService
	messageService port.MessageService
}

func NewHandler(socketService port.SocketService, roomService port.RoomService, messageService port.MessageService) Handler {
	return Handler{
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
	id, _ := socket.Session().Load("roomID")
	socketID, _ := socket.Session().Load("socketID")

	h.socketService.Delete(socketID.(string))
	//TODO: unsubscribe

	fmt.Println("OnClose", id, socketID)
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
	roomID, _ := socket.Session().Load("roomID")
	socketID, _ := socket.Session().Load("socketID")
	userID, _ := socket.Session().Load("userID")
	fmt.Println("OnMessage", roomID, socketID, message.Data, userID)

	u, _ := strconv.Atoi(userID.(string))

	msg := domain.Message{
		RoomID: roomID.(string),
		UserID: uint32(u),
		Body:   message.Data.String(),
	}
	_, err := h.messageService.Create(context.Background(), msg)
	if err != nil {
		fmt.Println(err)
	}

	subscribers := h.roomService.GetSubscribers(roomID.(string))
	for _, s := range subscribers {
		sid, _ := s.Session().Load("socketID")
		if sid != socketID {
			s.WriteMessage(message.Opcode, message.Bytes())
		}
	}
}
