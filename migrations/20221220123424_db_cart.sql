-- +goose Up
-- +goose StatementBegin
create table cart
(
    film_uuid uuid REFERENCES film (uuid) ON DELETE CASCADE,
    user_uuid uuid REFERENCES users (uuid) ON DELETE CASCADE,
    quantity int
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE cart;
-- +goose StatementEnd
