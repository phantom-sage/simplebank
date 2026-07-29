postgres18:
	docker container run -d -p5432:5432 --name postgres18 \
	-e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret \
	postgres:18-alpine

createdb:
	docker exec -it postgres18 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres18 dropdb simple_bank

migrateup:
	migrate -path db/migration \
	-database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database \
	"postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -cover -count=1 -v ./...

.PHONY: postgres18 createdb dropdb migrateup migratedown sqlc test