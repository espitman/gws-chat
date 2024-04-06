package socket

import (
	"fmt"
	"time"

	"github.com/espitman/gws-chat/chat-service/internal/core/port"
	"github.com/lxzan/gws"
)

const (
	PingInterval = 50 * time.Second
	PingWait     = 100 * time.Second
)

type Handler struct {
	socketService port.SocketService
	roomService   port.RoomService
}

func NewHandler(socketService port.SocketService, roomService port.RoomService) Handler {
	return Handler{
		socketService: socketService,
		roomService:   roomService,
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
	fmt.Println("OnMessage", roomID, socketID, message.Data)
	subscribers := h.roomService.GetSubscribers(roomID.(string))
	for _, s := range subscribers {
		sid, _ := s.Session().Load("socketID")
		if sid != socketID {
			s.WriteMessage(message.Opcode, message.Bytes())
		}
	}
}
