package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

func (s Server) ValidateInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		if err := s.validate.Struct(req); err != nil {
			s := status.New(codes.InvalidArgument, err.Error())
			return nil, s.Err()
		}
		return handler(ctx, req)
	}
}

func (s Server) AuthInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			s := status.New(codes.Unauthenticated, "Unauthenticated")
			return nil, s.Err()
		}
		token := md.Get("Authorization")
		if token == nil {
			s := status.New(codes.Unauthenticated, "Unauthenticated")
			return nil, s.Err()
		}
		userId := strings.Replace(token[0], "Bearer ", "", 1)
		if userId == "" {
			s := status.New(codes.Unauthenticated, "Unauthenticated")
			return nil, s.Err()
		}
		return handler(ctx, req)
	}
}
