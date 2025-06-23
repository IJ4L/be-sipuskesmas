gqlgen:
	go get github.com/99designs/gqlgen
	go run github.com/99designs/gqlgen generate
	go mod tidy

run:
	go run ./cmd/server.go

.PHONY: gqlgen run