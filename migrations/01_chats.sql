-- +goose Up
-- +goose StatementBegin
create table chats (
    chat_id     serial primary key,
    chat_name   varchar(25)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table chats;
-- +goose StatementEnd
