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