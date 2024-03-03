package api

import (
	"context"

	"chat-server/internal/converter"

	desc "chat-server/pkg/chat_api_v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// SendMessage get a new message in API layer.
func (i *Implementation) SendMessage(ctx context.Context, request *desc.SendMessageRequest) (*emptypb.Empty, error) {

	err := i.chatService.SendMessage(ctx, converter.ToSendMessageInput(request))
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to delete chat")
	}

	return &emptypb.Empty{}, nil
}
