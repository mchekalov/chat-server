package tests

import (
	"context"
	"fmt"
	serv "github.com/mchekalov/chat-server/internal/service/chat"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/mchekalov/chat-server/internal/model"
	"github.com/mchekalov/chat-server/internal/repository"
	repositoryMocks "github.com/mchekalov/chat-server/internal/repository/mocks"
)

func TestSaveMessage(t *testing.T) {
	t.Parallel()
	type chatRepositoryMockFunc func(mc *minimock.Controller) repository.ChatRepository

	type args struct {
		ctx context.Context
		req *model.Message
	}

	var (
		ctx   = context.Background()
		mc    = minimock.NewController(t)
		title = gofakeit.Animal()
		text  = gofakeit.Animal()
		id    = 0

		serviceErr = fmt.Errorf("repo error")

		req = &model.Message{
			ChatID:      int64(id),
			UserName:    title,
			MessageText: text,
		}
	)

	tests := []struct {
		name               string
		args               args
		err                error
		chatRepositoryMock chatRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: nil,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repositoryMocks.NewChatRepositoryMock(mc)
				mock.SaveMessageMock.Inspect(func(ctx context.Context, r *model.Message) {
					assert.Equal(mc, r.MessageText, req.MessageText)
					assert.Equal(mc, r.UserName, req.UserName)
				}).Return(nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: serviceErr,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repositoryMocks.NewChatRepositoryMock(mc)
				mock.SaveMessageMock.Expect(ctx, req).Return(serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatRepositoryMock := tt.chatRepositoryMock(mc)
			servChat := serv.NewService(chatRepositoryMock)

			err := servChat.SendMessage(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
		})
	}
}
