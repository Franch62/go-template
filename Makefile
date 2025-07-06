APP_NAME=myservice

.PHONY: build run lint test migrate

build:
	go build -o bin/$(APP_NAME) ./cmd/myservice

run:
	go run ./cmd/myservice

lint:
	golangci-lint run

test:
	go test -v ./...

migrate:
	go run scripts/migrate.go 
