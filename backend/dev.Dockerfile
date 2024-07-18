FROM golang:1.22

RUN go install github.com/air-verse/air@latest

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN chmod +x ./scripts/db_manage.sh
RUN chmod +x ./scripts/docker-entrypoint.sh

ENTRYPOINT ["./scripts/docker-entrypoint.sh"]

EXPOSE 8080
