-- +goose Up
CREATE TABLE users (
    id          UUID PRIMARY KEY, --Primary key
    created_at  TIMESTAMP NOT NULL, -- Not null
    updated_at  TIMESTAMP NOT NULL, -- NOt NUll
    name        VARCHAR(20) NOT NULL
);

-- +goose Down
DROP TABLE users;