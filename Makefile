run:
	go run cmd/api/main.go

dev:
	air

docker-up:
	docker compose up -d

docker-down:
	docker compose down

sonar:
	sonar-scanner \
		-Dsonar.projectKey=twonana-portfolio \
		-Dsonar.sources=. \
		-Dsonar.host.url=http://localhost:9000 \
		-Dsonar.token=sqp_d67de7edfd7b6c8a3d6b38ac2a4cf82ed1a5003e