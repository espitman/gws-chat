package grpc

import (
	"fmt"
	"github.com/espitman/gws-chat/message-service/internal/core/port"
	pb "github.com/espitman/gws-chat/pkg/protos/protogen/message-service"
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
	roomService port.RoomService,
	memberService port.MemberService,
	messageService port.MessageService,
	// +salvation NewServer
) *Server {
	handler := NewHandler(
		validate,
		roomService,
		memberService,
		messageService,
		// +salvation NewHandler

	)
	return &Server{
		port:     port,
		validate: validate,
		handler:  handler,
	}
}

func (s Server) registerServers(gs *grpc.Server) {
	pb.RegisterMessageServiceServer(gs, s.handler)
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
			s.AuthInterceptor(),
		),
	)
	s.registerServers(gs)
	reflection.Register(gs)
	if err := gs.Serve(listener); err != nil {
		panic(err)
	}
}
