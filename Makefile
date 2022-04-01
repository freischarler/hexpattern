SHELL:=/bin/bash -0 extglob
BINARY=main
VERSION=0.1.0
LDFLAGS=-ldflags "-X main.Version=${VERSION}"

#go tool commands
build:
	go build ${LDFLAGS} -i ${BINARY} main.go

run:
	@go run main.go

## docker compose
up: 
	docker-compose up --build
down:
	docker-compose down --remove-orphans
