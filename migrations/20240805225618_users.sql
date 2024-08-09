-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    user_type user_type NOT NULL,
    created_at TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
