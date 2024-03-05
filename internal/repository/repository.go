package repository

import (
	"context"

	"github.com/mchekalov/chat-server/internal/model"
)

// ChatRepository defines an interface for interacting with the repository layer
// to perform operations related to chat entities.
type ChatRepository interface {
	Create(ctx context.Context, chat *model.Chat) (int64, error)
	Delete(ctx context.Context, id *model.ChatDelete) error
	SaveMessage(ctx context.Context, message *model.Message) error
}
