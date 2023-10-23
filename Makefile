run:
	go run cmd/api/main.go

docker.up:
	docker compose up -d

docker.down:
	docker compose down

sonar:
	sonar-scanner \
		-Dsonar.projectKey=twonana-portfolio \
		-Dsonar.sources=. \
		-Dsonar.host.url=http://localhost:9000 \
		-Dsonar.token=sqp_d67de7edfd7b6c8a3d6b38ac2a4cf82ed1a5003e

nginx.run:
	docker run -d --name nginx-twonana -p 80:80 nginx:latest

nginx.copy:
	docker cp ./nginx/twonana.conf nginx-twonana:/etc/nginx/conf.d/