package domain

type Message struct {
	ID         uint32 `json:"ID"`
	RoomID     string `json:"roomID"`
	UserID     uint32 `json:"userID"`
	AudienceID uint32 `json:"audienceID"`
	Body       string `json:"body"`
	Time       string `json:"time"`
}
