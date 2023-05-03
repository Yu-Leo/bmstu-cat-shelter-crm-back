swag-init:
	swag init -g internal/endpoints/router.go
.PHONY: swag-init

d-run:
	docker build -t go-server .
	docker run --rm -p 9000:9000 go-server
.PHONY: d-run

run:
	go run ./...
.PHONY: run