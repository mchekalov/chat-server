package app

import (
	"context"
	"log"

	"github.com/mchekalov/chat-server/internal/api"
	"github.com/mchekalov/chat-server/internal/config"
	"github.com/mchekalov/chat-server/internal/repository"
	chatrepository "github.com/mchekalov/chat-server/internal/repository/chat"
	"github.com/mchekalov/chat-server/internal/service"
	chatservice "github.com/mchekalov/chat-server/internal/service/chat"
	"github.com/mchekalov/platform_common/pkg/closer"
	"github.com/mchekalov/platform_common/pkg/db"
	"github.com/mchekalov/platform_common/pkg/db/pg"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig
	dbClient   db.Client

	chatRepository repository.ChatRepository

	chatService service.ChatService

	chatImpl *api.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DsnString())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = chatrepository.NewRepository(s.DBClient(ctx))
	}

	return s.chatRepository
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatservice.NewService(
			s.ChatRepository(ctx))
	}

	return s.chatService
}

func (s *serviceProvider) ChatImpl(ctx context.Context) *api.Implementation {
	if s.chatImpl == nil {
		s.chatImpl = api.NewImplementation(s.ChatService(ctx))
	}

	return s.chatImpl
}
