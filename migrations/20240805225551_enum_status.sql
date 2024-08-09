-- +goose Up
-- +goose StatementBegin
CREATE TYPE status as enum ('created', 'approved', 'declined', 'on moderation');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE IF EXISTS status;
-- +goose StatementEnd
