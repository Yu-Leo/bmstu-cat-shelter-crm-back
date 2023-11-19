swag-init:
	swag init -g internal/endpoints/router.go
.PHONY: swag-init

d-run:
	docker build -t go-server .
	docker run --rm -p 9000:9000 -e APP_HOST='0.0.0.0' go-server
.PHONY: d-run

run:
	go run ./...
.PHONY: run

init-db:
	cat init/init.sql | sqlite3 database.db
.PHONY: init-db

build:
	go build -v ./...
.PHONY: build

gotools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2
	go install github.com/vektra/mockery/v2@v2.36.0
.PHONY: gotools

mocks:
	 go generate ./internal/repositories/gen.go
.PHONY: mocks

lint:
	golangci-lint run -v
.PHONY: lint

test:
	go test -v ./...
.PHONY: test

