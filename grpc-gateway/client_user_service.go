// This file generated automatically by gRPC gateway generator
// Generated at: 2024-04-10 01:16:08

package main

import (
	"log"

	userpb "github.com/espitman/gws-chat/pkg/protos/protogen/user-service"
	"google.golang.org/grpc"
)

func newUserServiceClient() userpb.UserServiceClient {
	cc, err := grpc.Dial("localhost:8001", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	return userpb.NewUserServiceClient(cc)
}
