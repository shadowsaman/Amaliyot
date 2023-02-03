

go:
	go run cmd/main.go

swag-init:
	swag init -g api/api.go -o api/docs

migration-up:
	migrate -path ./migrations/postgres/ -database 'postgres://samandar:saman107@localhost:5432/subscription_service?sslmode=disable' up

migration-down:
	migrate -path ./migrations/postgres/ -database 'postgres://samandar:saman107@localhost:5432/subscription_service?sslmode=disable' down
