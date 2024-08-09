-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS houses (
    id BIGINT PRIMARY KEY,
    address TEXT NOT NULL,
    year_built BIGINT NOT NULL,
    developer TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    last_flat_added TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS houses;
-- +goose StatementEnd
