-- +goose Up
-- +goose StatementBegin
CREATE TYPE user_type as enum ('client', 'moderator');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE IF EXISTS user_type;
-- +goose StatementEnd
