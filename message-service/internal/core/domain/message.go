package domain

import (
	"time"
)

type Message struct {
	ID     uint32
	RoomID string
	UserID uint32
	Body   string
	Time   time.Time
}
