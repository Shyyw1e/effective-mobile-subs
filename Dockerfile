# Используем официальный Go-образ для сборки
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/app

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/main .
COPY .env .env

EXPOSE 8080

CMD ["./main"]
