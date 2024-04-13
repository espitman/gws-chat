package port

import (
	"context"
)

type PusherService interface {
	Send(ctx context.Context, message []byte)
}
