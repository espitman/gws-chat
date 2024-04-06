package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/lxzan/gws"
)

const (
	PingInterval = 50 * time.Second
	PingWait     = 100 * time.Second
)

var upgrader *gws.Upgrader
var userSockets map[string]*gws.Conn
var subscribers map[string][]string

func init() {
	upgrader = gws.NewUpgrader(&Handler{}, &gws.ServerOption{
		ParallelEnabled:   true,                                 // Parallel message processing
		Recovery:          gws.Recovery,                         // Exception recovery
		PermessageDeflate: gws.PermessageDeflate{Enabled: true}, // Enable compression
	})
	userSockets = make(map[string]*gws.Conn)
	subscribers = make(map[string][]string)
}

func main() {
	router := httprouter.New()
	router.GET("/chat/:id", ChatHandler)
	fmt.Println("App run on port 6666")
	info()
	log.Fatal(http.ListenAndServe(":6666", router))

}

type Handler struct{}

func (c *Handler) OnOpen(socket *gws.Conn) {
	_ = socket.SetDeadline(time.Now().Add(PingInterval + PingWait))
	socketID, _ := socket.Session().Load("socketID")
	socket.WriteString(socketID.(string))
	fmt.Println("OnOpen")
}

func (c *Handler) OnClose(socket *gws.Conn, err error) {
	id, _ := socket.Session().Load("roomID")
	socketID, _ := socket.Session().Load("socketID")
	delete(userSockets, socketID.(string))

	var a []string
	for _, s := range subscribers[id.(string)] {
		if s != socketID {
			a = append(a, s)
		}
	}
	subscribers[id.(string)] = a

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
	for _, s := range subscribers[id.(string)] {
		if s != socketID {
			sk := userSockets[s]
			sk.WriteMessage(message.Opcode, message.Bytes())
		}
	}
}

func ChatHandler(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {

	auth := request.Header.Get("auth")
	id := ps.ByName("id")

	if auth != "saeed" {
		http.Error(writer, "Forbidden", 403)
	}

	socketID := uuid.NewString()

	//fmt.Println(auth, id)

	socket, err := upgrader.Upgrade(writer, request)

	if err != nil {
		return
	}

	socket.Session().Store("socketID", socketID)
	socket.Session().Store("roomID", id)
	userSockets[socketID] = socket
	subscribers[id] = append(subscribers[id], socketID)

	go func() {
		socket.ReadLoop() // Blocking prevents the context from being GC.
	}()

	//socket.WriteMessage(1, []byte(generateRandomString(8)))
	//rand := generateRandomString(8)

}

func info() {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("userSockets:", len(userSockets))
				for _, sk := range userSockets {
					id, _ := sk.Session().Load("roomID")
					socketID, _ := sk.Session().Load("socketID")
					fmt.Println(id, socketID)
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
