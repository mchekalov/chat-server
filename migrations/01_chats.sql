-- +goose Up
-- +goose StatementBegin
create table chats (
    chat_id     serial primary key,
    chat_name   text not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table chats;
-- +goose StatementEnd
