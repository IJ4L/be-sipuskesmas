-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email Text NOT NULL
);

-- +goose Down
DROP TABLE users;
