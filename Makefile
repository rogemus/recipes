.PHONY: test dev start dbUp dbDown

test:
	@echo "Testing..."
	go test ./test/... -v

dev:
	@echo "Starting dev server ..."
	air

start:
	@echo "Starting prod server ..."
	go run ./cmd/web/ -debug

dbUp:
	migrate -path=./db/migrations -database=${RECIPIES_DB_DSN} up

dbDown:
	migrate -path=./db/migrations -database=${RECIPIES_DB_DSN} down
