FROM golang:1.20.1-alpine3.17 AS builder
WORKDIR /dev
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

# run stage
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /dev/main .
COPY --from=builder /dev/migrate ./migrate
COPY /migration ./migration
COPY start.sh .
RUN chmod +x /app/start.sh
ENV DB_SOURCE=postgres://thopham:Lk8UZozxGk0BfPfAdfjsjUsptsYtMcKu@dpg-cgkogb8rddleudvr9ee0-a.singapore-postgres.render.com/nckhdb
EXPOSE 8080

ENTRYPOINT [ "/app/start.sh" ]
CMD [ "/app/main" ]