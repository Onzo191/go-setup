include .env

dev:
	air

run: swagger
	go run cmd/api/main.go

build: swagger
	go build -o bin/app cmd/api/main.go

## DOCKER INFRASTRUCTURE MANAGEMENT
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

## GENERATE API DOCUMENTATION
swagger:
	swag init -g cmd/api/main.go -d ./ --parseInternal


## MIGRATE
migrate:
	atlas migrate diff $(NAME) --env local

migrate-apply:
	atlas migrate apply --env local --url "$(DB_MIGRATE)"

migrate-status:
	atlas migrate status --env local --url "$(DB_MIGRATE)"

migrate-revert:
	atlas migrate down --env local --url "$(DB_MIGRATE)" 1