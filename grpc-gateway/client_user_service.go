// This file generated automatically by gRPC gateway generator
// Generated at: 2024-04-07 23:32:20

package main

import (
	userpb "github.com/espitman/gws-chat/pkg/protos/protogen/user-service"
	"google.golang.org/grpc"
	"log"
)

func newUserServiceClient() userpb.UserServiceClient {
	cc, err := grpc.Dial("localhost:8001", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	return userpb.NewUserServiceClient(cc)
}
