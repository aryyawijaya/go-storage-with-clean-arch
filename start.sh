#!/bin/sh

# script exit immidiately if any command return non-zero value
set -e

echo "run db migration"
migrate -path /app/db/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
# run all commands pass to this script
exec "$@"