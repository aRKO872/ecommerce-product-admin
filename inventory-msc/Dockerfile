# Stage 1: Build
FROM golang:1.24-alpine AS builder
WORKDIR /app
# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./
RUN go mod download
# Copy source
COPY . .
# Build binary
RUN go build -o inventory-srv-binary .
# Stage 2: Run
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/inventory-srv-binary .
EXPOSE 7003
CMD ["./inventory-srv-binary"]