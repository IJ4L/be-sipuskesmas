GOOSE_DIR=pkg/db/postgres/migrations
DB_URL=postgres://sipuskesmas:sipuskesmas@localhost:5432/sipuskesmas?sslmode=disable

gqlgen:
	go get github.com/99designs/gqlgen
	go run github.com/99designs/gqlgen generate
	go mod tidy

pg-sipuskesmas:
	docker run --name pg-sipuskesmas \
  	-e POSTGRES_USER=sipuskesmas \
  	-e POSTGRES_PASSWORD=sipuskesmas \
  	-e POSTGRES_DB=sipuskesmas \
  	-p 5432:5432 \
  	-d postgres:15

down-pg-sipuskesmas:
	docker stop pg-sipuskesmas && docker rm pg-sipuskesmas

create-migration:
	@read -p "Migration name: " name; \
	goose create "$$name" sql -dir $(GOOSE_DIR)

up-migrations:
	goose -dir $(GOOSE_DIR) postgres "$(DB_URL)" up

down-migrations:
	goose -dir $(GOOSE_DIR) postgres "$(DB_URL)" down

reset-db:
	goose -dir $(GOOSE_DIR) postgres "$(DB_URL)" down-to 0

status:
	goose -dir $(GOOSE_DIR) postgres "$(DB_URL)" status

run:
	go run ./cmd/server.go

.PHONY: gqlgen pg-sipuskesmas down-pg-sipuskesmas create-migrations up-migrations down-migrations reset-db status run