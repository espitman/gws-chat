// This file generated automatically by gRPC gateway generator
// Generated at: 2024-04-14 21:17:44

package main

import (
	messagepb "github.com/espitman/gws-chat/pkg/protos/protogen/message-service"
	"google.golang.org/grpc"
	"log"
)

func newMessageServiceClient() messagepb.MessageServiceClient {
	cc, err := grpc.Dial("localhost:8002", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	return messagepb.NewMessageServiceClient(cc)
}
