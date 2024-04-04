package socket

import (
	"fmt"
	"github.com/lxzan/gws"
	"time"
)

const (
	PingInterval = 50 * time.Second
	PingWait     = 100 * time.Second
)

type Handler struct {
}

func NewHandler() Handler {
	return Handler{}
}

func (c *Handler) OnOpen(socket *gws.Conn) {
	_ = socket.SetDeadline(time.Now().Add(PingInterval + PingWait))
	socketID, _ := socket.Session().Load("socketID")
	socket.WriteString(socketID.(string))
	fmt.Println("OnOpen")
}

func (c *Handler) OnClose(socket *gws.Conn, err error) {
	id, _ := socket.Session().Load("roomID")
	socketID, _ := socket.Session().Load("socketID")
	//delete(userSockets, socketID.(string))
	//var a []string
	//for _, s := range subscribers[id.(string)] {
	//	if s != socketID {
	//		a = append(a, s)
	//	}
	//}
	//subscribers[id.(string)] = a

	fmt.Println("OnClose", id, socketID)
}

func (c *Handler) OnPing(socket *gws.Conn, payload []byte) {
	_ = socket.SetDeadline(time.Now().Add(PingInterval + PingWait))
	_ = socket.WritePong(nil)
	fmt.Println("OnPing")
}

func (c *Handler) OnPong(socket *gws.Conn, payload []byte) {
	fmt.Println("OnPong")
}

func (c *Handler) OnMessage(socket *gws.Conn, message *gws.Message) {
	defer message.Close()
	id, _ := socket.Session().Load("roomID")
	socketID, _ := socket.Session().Load("socketID")
	fmt.Println("OnMessage", id, socketID)
	//for _, s := range subscribers[id.(string)] {
	//	if s != socketID {
	//		sk := userSockets[s]
	//		sk.WriteMessage(message.Opcode, message.Bytes())
	//	}
	//}
}
