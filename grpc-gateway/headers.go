package main

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
)

func getCtx(fctx fiberCtx) context.Context {
	reqHeaders := fctx.GetReqHeaders()
	//fmt.Println("userID:::", fctx.Locals("userID"))

	headers := make(map[string]string)
	if _, ok := reqHeaders["Authorization"]; ok {
		userID := fctx.Locals("userID").(string)
		headers["authorization"] = userID //reqHeaders["Authorization"][0]
	} else {
		headers["authorization"] = "Guest"
	}

	headers["trace-id"] = uuid.NewString()

	header := metadata.New(headers)
	ctx := metadata.NewOutgoingContext(fctx.Context(), header)
	return ctx
}
