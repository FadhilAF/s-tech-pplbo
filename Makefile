include .env

run:
	go run cmd/main.go

tidy:
	go mod tidy

migrate-up-postgres:
	migrate -path ./common/postgres/migration -database $(POSTGRES_CONNECTION_URL) -verbose up

migrate-down-postgres:
	migrate -path ./common/postgres/migration -database $(POSTGRES_CONNECTION_URL) -verbose down 

migrate-drop:
	migrate -path ./common/postgres/migration -database $(POSTGRES_CONNECTION_URL) -verbose drop

migrate-fresh: migrate-down migrate-up

sqlc:
	sqlc -f sqlc.yaml generate

.PHONY:	migrate-up migrate-down sqlc migrate-fresh run
