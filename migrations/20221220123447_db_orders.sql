-- +goose Up
-- +goose StatementBegin

create table orders
(
    uuid uuid NOT NULL DEFAULT uuid_generate_v4() primary key,
    film text,
    quantity int,
    user_uuid uuid REFERENCES users (uuid) ON DELETE CASCADE,
    date timestamp,
    status text
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE orders;
-- +goose StatementEnd
