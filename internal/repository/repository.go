package repository

import (
	"chat-server/internal/model"
	"context"
)

// ChatRepository defines an interface for interacting with the repository layer
// to perform operations related to chat entities.
type ChatRepository interface {
	Create(ctx context.Context, chat *model.Chat) (int64, error)
	Delete(ctx context.Context, id *model.ChatDelete) error
	SendMessage(ctx context.Context, message *model.Message) error
}
