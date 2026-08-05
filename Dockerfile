# Build stage.
FROM golang:1.26.5-alpine3.23 AS builder

WORKDIR /app

COPY . .

RUN go build -o main main.go

RUN apk add curl

RUN ARCH=$(uname -m) && \
    case "$ARCH" in \
      aarch64) ARCH="arm64" ;; \
      x86_64)  ARCH="amd64" ;; \
    esac && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.19.1/migrate.linux-${ARCH}.tar.gz | tar xvz

# Run stage.
FROM alpine:3.23

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate

COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration

EXPOSE 8080

CMD [ "/app/main" ]

ENTRYPOINT [ "/app/start.sh" ]