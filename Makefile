include .env

# templates
# 1. create new migration file
# - migrate create -ext sql -dir db/migration -seq <migration_name>

up-dev:
	docker compose -f compose.dev.yml up -d

down-dev:
	docker compose -f compose.dev.yml down

up-prod:
	docker compose -f compose.prod.yml up -d

down-prod:
	docker compose -f compose.prod.yml down

delete-image:
	docker rmi go-storage-with-clean-arch-api

logs-api:
	docker logs -f go-storage-api

migrateup-all:
	migrate \
		-path db/migration \
		-database ${DB_SOURCE_LOCAL} \
		-verbose \
		up

migrateup-1:
	migrate \
		-path db/migration \
		-database ${DB_SOURCE_LOCAL} \
		-verbose \
		up 1

migratedown-all:
	migrate \
		-path db/migration \
		-database ${DB_SOURCE_LOCAL} \
		-verbose \
		down

migratedown-1:
	migrate \
		-path db/migration \
		-database ${DB_SOURCE_LOCAL} \
		-verbose \
		down 1

sqlc:
	sqlc generate

mock-repo:
	mockgen \
		-package mockdb \
		-destination db/mock/all_repo.go \
		${MODULE_PATH}/util/test-helper Repo

mock-store:
	mockgen \
		-package mockdb \
		-destination db/mock/store.go \
		${MODULE_PATH}/db Store

query-update:
	make sqlc mock-repo mock-store

test:
	go test -timeout 30s -v -cover ./... -count=1

format:
	go fmt ./...

.PHONY:
	up-dev \
	down-dev \
	up-prod \
	down-prod \
	delete-image \
	logs-api \
	migrateup-all \
	migrateup-1 \
	migratedown-all \
	migratedown-1 \
	sqlc \
	mock \
	query-update \
	test \
	format \
