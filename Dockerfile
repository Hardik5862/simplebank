# Stage: build
FROM golang:1.24.3-alpine3.22 AS builder

WORKDIR /app

COPY . .

RUN go build -o main main.go

# Stage: run
FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/main .
COPY db/migration ./db/migration

EXPOSE 8005

CMD [ "/app/main" ]