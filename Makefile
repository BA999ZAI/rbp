.PHONY: gen-sqlc
gen-sqlc:
	docker run --rm -v .:/src -w /src sqlc/sqlc:1.24.0 -f ./internal/db/sqlc/sqlc.yml generate

.PHONY: build run test migrate clean
build:
	go build -o bin/api ./cmd/api

run:
	go run cmd/api/main.go

test:
	go test ./...

migrate:
	./scripts/migration.sh

clean:
	rm -rf bin