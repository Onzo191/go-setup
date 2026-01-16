run:
	go run cmd/api/main.go

dev:
	air

docker-infra-up:
	docker-compose -p go-infra \
	-f deployments/infra/postgres/docker-compose.yml \
	-f deployments/infra/redis/docker-compose.yml \
	up -d

docker-infra-down:
	docker-compose -p go-infra \
	-f deployments/infra/postgres/docker-compose.yml \
	-f deployments/infra/redis/docker-compose.yml \
	down