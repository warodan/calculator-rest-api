APP_NAME=calculator

include .env
export

.PHONY: run build test swag lint clean

run:
	go run ./cmd/app

build:
	go build -o bin/$(APP_NAME) ./cmd/app

test:
	go test ./... -v

swag:
	swag init --generalInfo cmd/app/main.go --output docs

clean:
	rm -rf bin/ docs/* $(APP_NAME)