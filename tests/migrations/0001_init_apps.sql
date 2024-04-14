-- +goose Up
INSERT INTO apps (name, secret)
VALUES ('test', 'test-secret')
ON CONFLICT (id)
DO NOTHING;

-- +goose Down
TRUNCATE apps RESTART IDENTITY;