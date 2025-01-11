# This is a template for creating a Dockerfile for a Go service. 
# Each Go microservise should contain a Dockerfile that is based on this template.

# Change <your_mod> to the name of the module you are working on
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o <your_mod>

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/<your_mod> .

CMD ["./<your_mod>"]
