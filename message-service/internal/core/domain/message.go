package domain

import (
	"time"

	messagepb "github.com/espitman/gws-chat/pkg/protos/protogen/message-service"
)

type Message struct {
	ID         uint32
	RoomID     string
	UserID     uint32
	UserName   string
	UserAvatar string
	Body       string
	Time       time.Time
	Type       messagepb.MessageType
}
