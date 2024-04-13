package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/espitman/gws-chat/chat-service/internal/core/domain"
	pushnotifications "github.com/pusher/push-notifications-go"
)

/**
 * pusherService implements port.pusherService interface
 */

const (
	instanceId = "4cd36c6b-4ddf-4a91-b59b-465c9480b538"
	secretKey  = "11F41A9A4B0648BB37FA45FDDDF1C15F0F7A2E3B358D958024EE8BB3D87963D8"
)

type PusherService struct {
}

func NewPusherService() *PusherService {
	return &PusherService{}
}

func (s *PusherService) Send(ctx context.Context, data []byte) {
	var message domain.Message
	err := json.Unmarshal(data, &message)
	if err != nil {
		fmt.Println("Error:", err)
	}

	beamsClient, err := pushnotifications.New(instanceId, secretKey)
	if err != nil {
		fmt.Println("Could not create Beams Client:", err.Error())
	}

	publishRequest := map[string]interface{}{
		"web": map[string]interface{}{
			"notification": map[string]interface{}{
				"title": "Hello",
				"body":  message.Body,
			},
		},
	}
	interest := "user-" + strconv.Itoa(int(message.AudienceID))
	pubId, err := beamsClient.PublishToInterests([]string{interest}, publishRequest)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Publish Id:", pubId, "interest:", interest)
	}
}
