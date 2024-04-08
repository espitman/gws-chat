package domain

type Room struct {
	ID     int
	RoomID string
	Users  string
}

type RoomInfo struct {
	ID       int
	RoomID   string
	User     User
	Audience User
}
