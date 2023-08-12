run:
	go run cmd/api/main.go

dev:
	air

swagger:
	swag init -g ./cmd/api/main.go -o ./docs

dockerup:
	docker compose up -d

dockerdown:
	docker compose down