package api

import (
	"context"

	desc "github.com/mchekalov/chat-server/pkg/chat_api_v1"

	"github.com/mchekalov/chat-server/internal/converter"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create creates a chat room in API layer
func (i *Implementation) Create(ctx context.Context, request *desc.CreateRequest) (*desc.CreateResponse, error) {
	output, err := i.chatService.Create(ctx, converter.ToCreateChatInput(request))
	if err != nil {

		return nil, status.Error(codes.Internal, "failed to create chat")
	}

	return converter.FromCreateChatInput(output), nil
}
