package http

import (
	"fmt"
	"net/http"
	"strconv"

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
	userService          port.UserService
}

func NewChatHandler(
	validate *validator.Validate,
	messageService port.MessageService,
	socketConnectService port.SocketConnetService,
	socketService port.SocketService,
	roomService port.RoomService,
	userService port.UserService,
) *ChatHandler {
	return &ChatHandler{
		validate:             validate,
		messageService:       messageService,
		socketConnectService: socketConnectService,
		socketService:        socketService,
		roomService:          roomService,
		userService:          userService,
	}
}

func (h *ChatHandler) ChatHandler(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	ctx := request.Context()
	queryValues := request.URL.Query()
	token := queryValues.Get("jwt")
	roomID := ps.ByName("id")
	if token == "" {
		http.Error(writer, "Forbidden", 403)
	}

	user, err2 := h.userService.Validate(ctx, token)
	if err2 != nil {
		http.Error(writer, "Forbidden", 403)
	}

	socket, err := h.socketConnectService.Open(writer, request, strconv.Itoa(int(user.ID)), roomID, token)
	if err != nil {
		fmt.Println(err)
	}

	go h.socketService.Save(socket)
	go h.roomService.Create(socket)
	go h.roomService.Subscribe(socket)

	go func() {
		socket.ReadLoop() // Blocking prevents the context from being GC.
	}()
}
