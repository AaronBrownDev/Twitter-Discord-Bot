FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install build dependencies for CGO
RUN apk add --no-cache gcc musl-dev

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o bot ./cmd/bot/main.go

# Use a smaller image for the final container
FROM alpine:latest

# SQLite dependencies
RUN apk add --no-cache \
    ca-certificates \
    sqlite \
    libc6-compat

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/bot .

# Command to run the executable
CMD ["./bot"]