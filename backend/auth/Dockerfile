# Dockerfile located at /backend/auth/Dockerfile

FROM golang:1.22-alpine as builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY backend/auth/go.mod backend/auth/go.sum ./

# Bring in the entire `backend` context so 'testing' is accessible
COPY backend /app/backend

WORKDIR /app/backend/auth

# Use go mod download inside the context that has all your modules
RUN go mod download

# Build the auth application
RUN go build -o auth-app .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/backend/auth/auth-app .

CMD ["./auth-app"]