package api

import (
	"github.com/mchekalov/chat-server/internal/service"
	desc "github.com/mchekalov/chat-server/pkg/chat_api_v1"
)

// Implementation represents the implementation of the chat API server.
type Implementation struct {
	desc.UnimplementedChatapiV1Server
	chatService service.ChatService
}

// NewImplementation creates a new instance of the chat API server implementation.
func NewImplementation(chatService service.ChatService) *Implementation {
	return &Implementation{
		chatService: chatService,
	}
}
