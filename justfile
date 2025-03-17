_default:
  just --list

# Run all tests
test:
	go test ./test/... -v

# Run api
run:
  go run ./api/cmd/api

# Run dev server with livereload
dev:
  watchexec --restart --debounce 2sec --exts go just run

# Run migration
db ACTION *N:
	migrate -path=./api/migrations -database=${RECIPES_DB_DSN} {{ACTION}} {{N}}

# Create new migration
dbCreate MIGRATION_FILE:
  migrate create -ext sql -dir api/migrations -seq {{MIGRATION_FILE}}
