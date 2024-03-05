package chat

import (
	"context"
	"fmt"

	"github.com/mchekalov/chat-server/internal/client/db"
	"github.com/mchekalov/chat-server/internal/model"
	"github.com/mchekalov/chat-server/internal/repository"
	"github.com/mchekalov/chat-server/internal/repository/chat/converter"

	"github.com/Masterminds/squirrel"
)

const (
	tableName      = "chats"
	chatnameColumn = "chat_name"
	chatIDColumn   = "chat_id"
)

const (
	tableMessages     = "messages"
	userNameColumn    = "user_name"
	messageTextColumn = "message_text"
)

type repo struct {
	db db.Client
	sq squirrel.StatementBuilderType
}

// NewRepository create new instance for repo object
func NewRepository(db db.Client) repository.ChatRepository {
	return &repo{
		db: db,
		sq: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)}
}

func (r *repo) Create(ctx context.Context, chat *model.Chat) (int64, error) {
	// Better place to convert from service model to repo model
	chatRepo := converter.FromChatToRepo(chat)

	builder := r.sq.Insert(tableName).
		Columns(chatnameColumn).
		Values(chatRepo.ChatName).
		Suffix(fmt.Sprintf("RETURNING %v", chatIDColumn))

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "chat_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Delete(ctx context.Context, id *model.ChatDelete) error {
	builder := r.sq.Delete(tableName).
		Where(squirrel.Eq{chatIDColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chat_repository.Delete",
		QueryRaw: query,
	}

	r.db.DB().QueryRowContext(ctx, q, args...)

	return nil
}

func (r *repo) SaveMessage(ctx context.Context, message *model.Message) error {
	builder := r.sq.Insert(tableMessages).
		Columns(chatIDColumn, userNameColumn, messageTextColumn).
		Values(message.ChatID, message.UserName, message.MessageText)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chat_repository.SendMessage",
		QueryRaw: query,
	}

	r.db.DB().QueryRowContext(ctx, q, args)

	return nil
}
