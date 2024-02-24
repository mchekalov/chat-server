-- +goose Up
-- +goose StatementBegin
create table messages (
    message_id      serial primary key,
    chat_id         int,
    user_name       text not null,
    message_text    text not null,
    message_created_at timestamp not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table messages;
-- +goose StatementEnd