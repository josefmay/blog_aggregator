-- +goose Up
CREATE TABLE feeds (
    id          UUID PRIMARY KEY, --Primary key
    created_at  TIMESTAMP NOT NULL, -- Not null
    updated_at  TIMESTAMP NOT NULL, -- NOt NUll
    name        TEXT NOT NULL,
    url         TEXT UNIQUE NOT NULL,
    user_id     UUID NOT NULL REFERENCES users
                            ON DELETE CASCADE,
    FOREIGN KEY(user_id) REFERENCES users (id)
);

-- +goose Down
DROP TABLE feeds;