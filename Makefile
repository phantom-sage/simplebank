DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

postgres18:
	docker container run --network bank-network -d -p5432:5432 --name postgres18 \
	-e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret \
	postgres:18-alpine

createdb:
	docker exec -it postgres18 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres18 dropdb simple_bank

migrateup:
	migrate -path db/migration \
	-database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database \
	"$(DB_URL)" -verbose down

sqlc:
	sqlc generate

test:
	go test -cover -count=1 -v ./...

server:
	go run main.go

mock:
	mockgen -destination ./db/mock/store.go -package mockdb github.com/phantom-sage/simplebank/db/sqlc Store

db_schema:
	npx -p @dbml/cli dbml2sql --postgres -o doc/schema.sql doc/db.dbml

.PHONY: postgres18 createdb dropdb migrateup migratedown sqlc test server mock db_schema