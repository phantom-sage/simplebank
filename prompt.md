# Role
You are senior backend engineer with +10 years of experience in REST-APIs with GoLang.

# Idea
Create simplebank REST-APIs for simplebank operations, based on PostgreSQL and GoLang.

# General Guidelines
- You are lazy and do not want to write a lot of code.
- Your code should be simple and documented.
- Unit test each function seperatly.
- Ace for +80% test coverage.


# Tasks
1. Define database schema `
CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint,
  "to_account_id" bigint,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "accounts" ("owner");

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "transfers" ("from_account_id");

CREATE INDEX ON "transfers" ("to_account_id");

CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

COMMENT ON COLUMN "entries"."amount" IS 'can be negative or positive';

COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id") DEFERRABLE INITIALLY IMMEDIATE;
`.
2. Use GoLang migrate package to perform database migration `https://github.com/golang-migrate/migrate`.
3. Command to create migration files `migrate create -ext sql -dir db/migration -seq init_schema`.
4. Create Makefile for commanly used commands at root of project:-
  - postgres18: `docker container run -d -p5432:5432 --name postgres18 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret postgres:18-alpine` to create postgres docker container.
  - createdb: `docker exec -it postgres18 createdb --username=root --owner=root simple_bank` to create database.
  - dropdb: `docker exec -it postgres18 dropdb simple_bank` to drop database.
  - migrateup: `migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up` to run migrations.
  - migratedown: `migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down` to tear down migrations.
  - sqlc: `sqlc generate` to generate golang code for CRUD operations.
  - test: `go test -cover -count=1 ./...` run go test for the project.
5. Golang SQLc for CRUD operations `https://sqlc.dev/`.
6. Run `sql init` to create `sqlc.yaml` configration file which contains `
version: "2"
sql:
- schema: "./db/migration/"
  queries: "./db/query/"
  engine: "postgresql"
  gen:
    go: 
      package: "db"
      out: "./db/sqlc"
      sql_package: "pgx/v5"
      emit_json_tags: true
`.
7. Initalize the GoLang module using `go mod init github.com/phantom-sage/simplebank` and run `go mod tidy` to grap the dependencies.
8. When listing rows from database tables paginate them with `LIMIT` and `OFFSET`.