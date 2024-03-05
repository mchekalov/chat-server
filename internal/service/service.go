package service

import (
	"context"

	"github.com/mchekalov/chat-server/internal/model"
)

// ChatService defines an interface for interacting with the service layer
// to perform operations related to chat entities.
type ChatService interface {
	Create(ctx context.Context, chat *model.Chat) (int64, error)
	Delete(ctx context.Context, id *model.ChatDelete) error
	SendMessage(ctx context.Context, message *model.Message) error
}
