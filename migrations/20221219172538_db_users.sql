
-- +goose Up
-- +goose StatementBegin
create table users
(
    uuid uuid NOT NULL DEFAULT uuid_generate_v4() primary key,
    name text,
    role int,
    pass text
);



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd