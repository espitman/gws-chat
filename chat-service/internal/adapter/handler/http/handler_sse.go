package http

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/espitman/gws-chat/chat-service/internal/core/port"
	"github.com/julienschmidt/httprouter"
	"github.com/r3labs/sse/v2"
)

type SSEHandler struct {
	sseServer   *sse.Server
	pubSub      *gochannel.GoChannel
	userService port.UserService
}

func NewSSEHandler(sseServer *sse.Server, pubSub *gochannel.GoChannel, userService port.UserService) *SSEHandler {
	return &SSEHandler{
		sseServer:   sseServer,
		pubSub:      pubSub,
		userService: userService,
	}
}

func (h *SSEHandler) setHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Type")

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
}

func (h *SSEHandler) SSEHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	h.setHeaders(w)

	userID := ps.ByName("id")
	uid, _ := strconv.Atoi(userID)

	go h.userService.SetOnline(context.TODO(), uint32(uid), true)

	ctx := r.Context()
	messages, _ := h.pubSub.Subscribe(ctx, userID)

	for msg := range messages {
		fmt.Fprintf(w, "data: %s\n\n", fmt.Sprintf("%s", msg.Payload))
		w.(http.Flusher).Flush()
		msg.Ack()
	}

	for {
		select {
		case <-ctx.Done():
			go h.userService.SetOnline(context.TODO(), uint32(uid), false)
			fmt.Println("left")
			ctx.Done()
			return
		default:
		}
	}
}
