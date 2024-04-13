package subscriber

import (
	"context"
	"fmt"

	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/espitman/gws-chat/chat-service/internal/core/port"
)

type ChatMessageSubscriber struct {
	pubSub        *gochannel.GoChannel
	pusherService port.PusherService
}

func NewChatMessageSubscriber(
	pubSub *gochannel.GoChannel,
	pusherService port.PusherService,
) *ChatMessageSubscriber {
	chatMessageSubscriber := ChatMessageSubscriber{
		pubSub:        pubSub,
		pusherService: pusherService,
	}
	return &chatMessageSubscriber
}

func (s *ChatMessageSubscriber) Subscribe() {
	ctx := context.TODO()
	messages, err := s.pubSub.Subscribe(ctx, "chat-message")
	if err != nil {
		fmt.Println(err)
	}
	for msg := range messages {
		fmt.Printf("chat-message:::received message: %s, payload: %s\n", msg.UUID, string(msg.Payload))
		s.pusherService.Send(ctx, msg.Payload)
		msg.Ack()
	}
}
