package main

import (
	"context"
	"log"
	"net"

	desc "chat-server/pkg/chat_api_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	desc.UnimplementedChatapiV1Server
}

func (s *server) Create(_ context.Context, r *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Create new chat: %v", r.GetUsernames())
	return &desc.CreateResponse{
		Id: 4,
	}, nil
}

func (s *server) Delete(_ context.Context, r *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Delete chat by Id: %v", r.GetId())
	return new(emptypb.Empty), nil
}

func (s *server) SendMessage(_ context.Context, r *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("Send message: %v", r.GetInfo())
	return new(emptypb.Empty), nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatapiV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
