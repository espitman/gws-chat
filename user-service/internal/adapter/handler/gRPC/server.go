package grpc

import (
	"fmt"
	pb "github.com/espitman/gws-chat/pkg/protos/protogen/order-service"
	"github.com/espitman/gws-chat/user-service/internal/core/port"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Server struct {
	port     string
	validate *validator.Validate
	handler  *Handler
}

func NewServer(
	port string,
	validate *validator.Validate,
	userService port.UserService,
	// +salvation NewServer
) *Server {
	handler := NewHandler(
		validate,
		userService,
		// +salvation NewHandler
	)
	return &Server{
		port:     port,
		validate: validate,
		handler:  handler,
	}
}

func (s Server) registerServers(gs *grpc.Server) {
	pb.RegisterOrderServiceServer(gs, s.handler)
}

func (s Server) Run() {
	listener, err := net.Listen("tcp", "localhost:"+s.port)
	if err != nil {
		panic(err)
	}
	fmt.Println("starting server:" + s.port)
	gs := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			s.ValidateInterceptor(),
			//s.AuthInterceptor(),
		),
	)
	s.registerServers(gs)
	reflection.Register(gs)
	if err := gs.Serve(listener); err != nil {
		panic(err)
	}
}
