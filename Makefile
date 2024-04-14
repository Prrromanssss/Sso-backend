migrate-up:
	goose postgres "postgres://postgres:6,62607004@localhost:5434/sso?sslmode=disable" up
migrate-down:
	goose postgres "postgres://postgres:6,62607004@localhost:5434/sso?sslmode=disable" down
migrate-stat:
	goose postgres "postgres://postgres:6,62607004@localhost:5434/sso?sslmode=disable" status
