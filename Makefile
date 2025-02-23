.PHONY: test dev start dbUp dbDown

test:
	@echo "Testing..."
	go test ./test/... -v

dev:
	@echo "Starting dev server ..."
	air

start:
	@echo "Starting prod server ..."
	go run ./cmd/web/

dbUp:
	migrate -path db/migrations -database sqlite3://recipies.db up

dbDown:
	migrate -path db/migrations -database sqlite3://recipies.db down
