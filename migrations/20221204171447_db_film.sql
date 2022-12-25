-- +goose Up
-- +goose StatementBegin
create table film
(
    uuid        uuid NOT NULL DEFAULT uuid_generate_v4() primary key,
    name        text,
    release     int,
    grade       float,
    genre       text,
    Price       int,
    whatch_time int,
    summary     text,
    image       text,
    video       text
);



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE films;
-- +goose StatementEnd
