package chat

import (
	"chat-server/internal/model"
	repository "chat-server/internal/repository"
	"chat-server/internal/service"
	"context"
)

type serv struct {
	chatRepository repository.ChatRepository
}

// NewService creates instance of service layer
func NewService(chatRepository repository.ChatRepository) service.ChatService {
	return &serv{chatRepository: chatRepository}
}

func (s *serv) Create(ctx context.Context, in *model.Chat) (int64, error) {

	output, err := s.chatRepository.Create(ctx, in)
	if err != nil {
		return 0, err
	}

	return output, nil
}

func (s *serv) Delete(ctx context.Context, in *model.ChatDelete) error {

	err := s.chatRepository.Delete(ctx, in)
	if err != nil {
		return err
	}

	return nil

}

func (s *serv) SendMessage(ctx context.Context, in *model.Message) error {

	err := s.chatRepository.SendMessage(ctx, in)
	if err != nil {
		return err
	}

	return nil

}
