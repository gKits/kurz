FROM golang:1.24-alpine AS build

WORKDIR /data

COPY . .

RUN "go build -o /bin/server ./cmd/server/main.go"
RUN "go build -o /bin/migrate ./cmd/migrate/main.go"

FROM alpine:latest AS run

COPY entrypoint.sh .

COPY --from=build --chown=1000:1000 /bin/server /bin/server
COPY --from=build --chown=1000:1000 /bin/migrate /bin/migrate

ENTRYPOINT [ "sh", "entrypoint.sh" ]
