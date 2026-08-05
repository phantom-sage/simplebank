#!/bin/sh

set -e

echo "run db migration"
/app/migrate -path /app/migration -database "postgresql://root:secret@postgres:5432/simple_bank?sslmode=disable" -verbose up

echo "start the app"
exec "$@"