package http

import (
	"net/http"

	"github.com/espitman/gws-chat/chat-service/internal/core/port"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

type ChatHandler struct {
	validate             *validator.Validate
	messageService       port.MessageService
	socketConnectService port.SocketConnetService
	socketService        port.SocketService
	roomService          port.RoomService
}

func NewChatHandler(
	validate *validator.Validate,
	messageService port.MessageService,
	socketConnectService port.SocketConnetService,
	socketService port.SocketService,
	roomService port.RoomService,
) *ChatHandler {
	return &ChatHandler{
		validate:             validate,
		messageService:       messageService,
		socketConnectService: socketConnectService,
		socketService:        socketService,
		roomService:          roomService,
	}
}

func (h *ChatHandler) ChatHandler(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {

	auth := "" //request.Header.Get("auth")
	id := ps.ByName("id")
	//if auth != "saeed" {
	//	http.Error(writer, "Forbidden", 403)
	//}

	socket, err := h.socketConnectService.Open(writer, request, auth, id)
	if err != nil {
		return
	}

	h.socketService.Save(socket)
	h.roomService.Create(socket)
	h.roomService.Subscribe(socket)

	go func() {
		socket.ReadLoop() // Blocking prevents the context from being GC.
	}()

	//socket.WriteMessage(1, []byte(generateRandomString(8)))
	//rand := generateRandomString(8)

}
