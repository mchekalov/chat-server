package tests

import (
	"context"
	serv "github.com/mchekalov/chat-server/internal/service/chat"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/mchekalov/chat-server/internal/model"
	"github.com/mchekalov/chat-server/internal/repository"
	repositoryMocks "github.com/mchekalov/chat-server/internal/repository/mocks"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type chatRepositoryMockFunc func(mc *minimock.Controller) repository.ChatRepository

	type args struct {
		ctx context.Context
		req *model.Chat
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id    = gofakeit.Int64()
		title = gofakeit.Animal()

		serviceErr = status.Error(codes.Internal, "failed to create chat")

		req = &model.Chat{
			ChatName: title,
		}
	)

	tests := []struct {
		name               string
		args               args
		want               int64
		err                error
		chatRepositoryMock chatRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: id,
			err:  nil,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repositoryMocks.NewChatRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, req).Return(id, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: 0,
			err:  serviceErr,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repositoryMocks.NewChatRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, req).Return(0, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatRepositoryMock := tt.chatRepositoryMock(mc)
			servChat := serv.NewService(chatRepositoryMock)

			newID, err := servChat.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
