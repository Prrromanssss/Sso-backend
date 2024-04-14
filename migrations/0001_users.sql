-- +goose Up
CREATE TABLE IF NOT EXISTS users
(
    id        int GENERATED ALWAYS AS IDENTITY,
    email     text NOT NULL UNIQUE,
    pass_hash oid NOT NULL,

    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE IF EXISTS users;