-- +goose Up
CREATE TABLE IF NOT EXISTS apps
(
    id        int GENERATED ALWAYS AS IDENTITY,
    name      text NOT NULL UNIQUE,
    secret    text NOT NULL UNIQUE,

    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE IF EXISTS apps;