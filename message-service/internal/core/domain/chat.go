package domain

import (
	"time"
)

type Chat struct {
	User   User
	RoomID string
	Body   string
	Time   time.Time
}
