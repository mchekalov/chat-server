package api

import (
	"context"

	"github.com/mchekalov/chat-server/internal/converter"
	desc "github.com/mchekalov/chat-server/pkg/chat_api_v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Delete deletes a chat room in API layer
func (i *Implementation) Delete(ctx context.Context, request *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := i.chatService.Delete(ctx, converter.ToDeleteChatInput(request))
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to delete chat")
	}

	return &emptypb.Empty{}, nil
}
