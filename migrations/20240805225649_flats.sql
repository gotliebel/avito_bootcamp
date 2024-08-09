-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS flats (
    house_id BIGINT NOT NULL,
    number BIGINT NOT NULL,
    price BIGINT NOT NULL,
    rooms INT NOT NULL,
    status status NOT NULL,
    created_at TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS flats;
-- +goose StatementEnd
