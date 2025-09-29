# Stage 1: Build
FROM golang:1.24-alpine AS builder
WORKDIR /app
# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./
RUN go mod download
# Copy source
COPY . .
# Build binary
RUN go build -o orders-msc-binary .
# Stage 2: Run
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/orders-msc-binary .
EXPOSE 7004
CMD ["./orders-msc-binary"]