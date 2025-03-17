default:
  just --list

# Run all tests
test:
	go test ./test/... -v

# Starting dev server
dev:
	go run ./cmd/web/ -debug

# Starting dev server with live reloard
serv:
  watchexec --restart --debounce 2sec --exts go,tmpl just dev

# Starting prod server
start:
	go run ./cmd/web/

# Run migration
db ACTION *N:
	migrate -path=./db/migrations -database=${RECIPES_DB_DSN} {{ACTION}} {{N}}

# Create new migration
migration NAME:
  migrate create -ext sql -dir db/migrations -seq {{NAME}}
