package domain

import "github.com/lxzan/gws"

type Subscriber struct {
	ID     string
	RoomID string
	Socket *gws.Conn
}
