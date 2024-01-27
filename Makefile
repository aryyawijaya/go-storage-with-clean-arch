include .env

# templates
# 1. create new migration file
# - migrate create -ext sql -dir db/migration -seq <migration_name>

compose-up-dev:
	docker compose -f compose.dev.yml up -d

compose-down-dev:
	docker compose -f compose.dev.yml down

compose-up-prod:
	docker compose -f compose.prod.yml up -d

compose-down-prod:
	docker compose -f compose.prod.yml down

delete-image:
	docker rmi go-storage-api

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
	compose-up-dev \
	compose-down-dev \
	compose-up-prod \
	compose-down-prod \
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
