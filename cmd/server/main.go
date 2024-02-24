package main

import (
	"context"
	"flag"
	"log"
	"net"

	"chat-server/config"
	desc "chat-server/pkg/chat_api_v1"

	"github.com/brianvoe/gofakeit"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	desc.UnimplementedChatapiV1Server
	pgx *pgx.Conn
}

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func (s *server) Create(ctx context.Context, r *desc.CreateRequest) (*desc.CreateResponse, error) {

	log.Printf("Create new chat: %v", r.GetUsernames())

	row := s.pgx.QueryRow(ctx, "INSERT INTO chats (chat_name) VALUES ($1) RETURNING chat_id", r.Chatname)
	var chatID, usID int64
	err := row.Scan(&chatID)

	for _, v := range r.GetUsernames() {
		row := s.pgx.QueryRow(ctx, "INSERT INTO users (user_name, chat_id) VALUES ($1, $2) RETURNING user_id",
			v, chatID)
		err = row.Scan(&usID)
	}

	return &desc.CreateResponse{
		Id: chatID,
	}, err
}

func (s *server) Delete(ctx context.Context, r *desc.DeleteRequest) (*emptypb.Empty, error) {

	_, err := s.pgx.Exec(ctx, "DELETE FROM chats WHERE chat_id=$1", r.Id)

	return new(emptypb.Empty), err
}

func (s *server) SendMessage(ctx context.Context, r *desc.SendMessageRequest) (*emptypb.Empty, error) {

	row := s.pgx.QueryRow(ctx, "INSERT INTO messages (chat_id, user_name, message_text) VALUES ($1, $2, $3) RETURNING message_id",
		gofakeit.Uint8(), r.Info.From, r.Info.Text)
	var messageID int64
	err := row.Scan(&messageID)

	return new(emptypb.Empty), err
}

func main() {

	ctx := context.Background()
	flag.Parse()

	err := config.Load(configPath)
	if err != nil {
		log.Printf("Failed to read ENV file: %v", err)
	}

	grpcConf, err := config.NewGRPCConfig()
	if err != nil {
		log.Printf("Failed to get gRPC config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConf.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// get postgres db connect

	pgConfig, err := config.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to load postgres config: %v", err)
	}

	con, err := pgx.Connect(ctx, pgConfig.DsnString())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer func() {
		if err = con.Close(ctx); err != nil {
			log.Printf("Error when closing connection: %v", err)
		}
	}()

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatapiV1Server(s, &server{pgx: con})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
