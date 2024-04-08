package domain

type Message struct {
	ID     uint32
	RoomID string
	UserID uint32
	Body   string
	Time   string
}
