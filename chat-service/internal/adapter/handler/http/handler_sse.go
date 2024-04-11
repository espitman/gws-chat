package http

import (
	"fmt"
	"net/http"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/julienschmidt/httprouter"
	"github.com/r3labs/sse/v2"
)

type SSEHandler struct {
	sseServer *sse.Server
	pbuSub    *gochannel.GoChannel
}

func NewSSEHandler(sseServer *sse.Server, pbuSub *gochannel.GoChannel) *SSEHandler {
	return &SSEHandler{
		sseServer: sseServer,
		pbuSub:    pbuSub,
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
	fmt.Println("userID..", userID)

	ctx := r.Context()
	messages, _ := h.pbuSub.Subscribe(ctx, userID)
	//go process(messages)

	for msg := range messages {
		fmt.Printf("received message: %s, payload: %s\n", msg.UUID, string(msg.Payload))
		//fmt.Fprintf(w, string(msg.Payload))
		fmt.Fprintf(w, "data: %s\n\n", fmt.Sprintf("%s", msg.Payload))
		w.(http.Flusher).Flush()
		msg.Ack()
	}

	//for {
	//	select {
	//	case <-ctx.Done():
	//		fmt.Println("left")
	//		ctx.Done()
	//		return
	//	default:
	//		i++
	//		fmt.Println(i)
	//		fmt.Fprintf(w, "data: %s\n\n", fmt.Sprintf("Event %d", i))
	//		time.Sleep(2 * time.Second)
	//		w.(http.Flusher).Flush()
	//	}
	//}
}

func process(messages <-chan *message.Message) {
	for msg := range messages {
		fmt.Printf("received message: %s, payload: %s\n", msg.UUID, string(msg.Payload))

		// we need to Acknowledge that we received and processed the message,
		// otherwise, it will be resent over and over again.
		msg.Ack()
		fmt.Println("ACK")
	}
}
