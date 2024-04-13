package subsctib

import (
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/espitman/gws-chat/chat-service/internal/adapter/subscriber"
	"github.com/espitman/gws-chat/chat-service/internal/core/port"
)

func Run(pubSub *gochannel.GoChannel, pusherService port.PusherService) {
	chatMessageSubscriber := subscriber.NewChatMessageSubscriber(pubSub, pusherService)
	chatMessageSubscriber.Subscribe()
}
