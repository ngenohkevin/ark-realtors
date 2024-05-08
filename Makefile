#! include .env
network:
	@docker network create ${NETWORK_NAME}

up:
	@docker compose up

down:
	@docker compose down -v --remove-orphans

background:
	@docker compose up -d --remove-orphans

start_container:
	@docker start ${DB_CONTAINER_NAME}

stop_container:
	@docker stop ${DB_CONTAINER_NAME}

rm_container:
	@docker rm ${DB_CONTAINER_NAME}

postgres:
	@docker run --name ${DB_CONTAINER_NAME} -e POSTGRES_USER=${USER} -e POSTGRES_PASSWORD=${DB_PASSWORD} -p ${DB_PORT}:5432 -d ${DB_IMAGE}

createdb:
	@docker exec -it ${DB_CONTAINER_NAME} createdb --username=${USER} --owner=${USER} ${DB_NAME}

dropdb:
	@docker exec -it ${DB_CONTAINER_NAME} dropdb ${DB_NAME}

migration:
	@migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	@migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	@migrate -path db/migration -database "$(DB_URL)" -verbose down

migrateSession:
	@migrate create -ext sql -dir db/migration -seq add_sessions

migrateup1:
	@migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown1:
	@migrate -path db/migration -database "$(DB_URL)" -verbose down 1

migratePictures:
	@migrate create -ext sql -dir db/migration -seq add_pictures

migrateup2:
	@migrate -path db/migration -database "$(DB_URL)" -verbose up 2

migratedown2:
	@migrate -path db/migration -database "$(DB_URL)" -verbose down 2

sqlc:
	@sqlc generate

mock:
	@mockgen -package mockdb -destination pkg/mock/store.go github.com/ngenohkevin/ark-realtors/internal/store Store

test:
	@go test -v -cover ./...


.PHONY: network run stop postgres createdb dropdb migration migrateup migratedown rm_container stop_container sqlc mock test migrateSession migrateup1 migratedown1 migratePictures migrateup2 migratedown2