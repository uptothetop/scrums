FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o user-service

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/user-service .

CMD ["./user-service"]
