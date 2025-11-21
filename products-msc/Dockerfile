# Stage 1: Build
FROM golang:1.24-bullseye AS builder

# Install C build tools and librdkafka headers
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
      build-essential pkg-config librdkafka-dev git ca-certificates && \
    rm -rf /var/lib/apt/lists/*

ENV CGO_ENABLED=1
WORKDIR /app
# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./
RUN go mod download
# Copy source
COPY . .
# Build binary
RUN go build -o products-msc-binary .
# Stage 2: Run
FROM debian:bullseye-slim
# install librdkafka runtime and ca-certs
RUN apt-get update && \
    apt-get install -y --no-install-recommends librdkafka1 ca-certificates && \
    rm -rf /var/lib/apt/lists/*
WORKDIR /root/
COPY --from=builder /app/products-msc-binary .
EXPOSE 7005
CMD ["./products-msc-binary"]