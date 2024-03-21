package tests

import (
	"context"
	"github.com/mchekalov/chat-server/internal/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/mchekalov/chat-server/internal/model"
	"github.com/mchekalov/chat-server/internal/service"
	serviceMocks "github.com/mchekalov/chat-server/internal/service/mocks"
	desc "github.com/mchekalov/chat-server/pkg/chat_api_v1"
)

func TestSendMessage(t *testing.T) {
	t.Parallel()
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *desc.SendMessageRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id    int64 = 0
		title       = gofakeit.Animal()
		text        = gofakeit.Animal()
		time        = timestamppb.Now()

		serviceErr = status.Error(codes.Internal, "failed to delete chat")

		req = &desc.SendMessageRequest{
			Info: &desc.MessageWrap{
				From:      title,
				Text:      text,
				Timestamp: time,
			},
		}

		info = &model.Message{
			ChatID:      id,
			UserName:    title,
			MessageText: text,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *emptypb.Empty
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: &emptypb.Empty{},
			err:  nil,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.SendMessageMock.Expect(ctx, info).Return(nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.SendMessageMock.Expect(ctx, info).Return(serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatServiceMock := tt.chatServiceMock(mc)
			apiChat := api.NewImplementation(chatServiceMock)

			newID, err := apiChat.SendMessage(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
