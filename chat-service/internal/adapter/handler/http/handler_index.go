package http

import (
	"fmt"
	"net/http"

	"github.com/espitman/gws-chat/chat-service/internal/core/domain"
	"github.com/espitman/gws-chat/chat-service/internal/core/port"
	"github.com/julienschmidt/httprouter"
)

type IndexHandler struct {
	roomService port.RoomService
}

type IndexHandlerDto struct {
	rooms       []domain.Room
	subscribers []domain.Subscriber
}

func NewIndexHandler(roomService port.RoomService) *IndexHandler {
	return &IndexHandler{roomService: roomService}
}

func (h *IndexHandler) IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rooms := h.roomService.GetAllRooms()
	subscribers := h.roomService.GetAllSubscribers()
	resp := IndexHandlerDto{
		rooms:       rooms,
		subscribers: subscribers,
	}
	_, _ = fmt.Fprint(w, resp)
}
