package socket

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
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
	//token, _ := socket.Session().Load("token")
	fmt.Println("OnMessage", roomID, socketID, message.Data, userID)

	u, _ := strconv.Atoi(userID.(string))
	audienceID, err := h.roomService.GetAudience(context.TODO(), roomID.(string), uint32(u))
	if err != nil {
		fmt.Println(err)
	}

	msg := domain.Message{
		RoomID:     roomID.(string),
		UserID:     uint32(u),
		Body:       message.Data.String(),
		AudienceID: audienceID,
	}
	_, err = h.messageService.Create(context.Background(), msg)
	if err != nil {
		fmt.Println(err)
	}

	go h.publishMessages(strconv.Itoa(int(audienceID)), msg)

	subscribers := h.roomService.GetSubscribers(roomID.(string))
	for _, s := range subscribers {
		sid, _ := s.Session().Load("socketID")
		if sid != socketID {
			s.WriteMessage(message.Opcode, message.Bytes())
		}
	}
}

func (h *Handler) publishMessages(userID string, data domain.Message) {
	go h.publishSSE(userID, data)
	go h.publishPush(userID, data)
}

func (h *Handler) publishSSE(userID string, data domain.Message) {
	msg := message.NewMessage(watermill.NewUUID(), message.Payload(data.Body))
	if err := h.pubSub.Publish(userID, msg); err != nil {
		fmt.Println(err)
	}
}

func (h *Handler) publishPush(userID string, data domain.Message) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	msg := message.NewMessage(watermill.NewUUID(), jsonData)
	if err := h.pubSub.Publish("chat-message", msg); err != nil {
		fmt.Println(err)
	}
}
