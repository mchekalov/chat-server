-- +goose Up
-- +goose StatementBegin
create table users (
    user_id     serial primary key,
    user_name   text not null,
    chat_id     int
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd