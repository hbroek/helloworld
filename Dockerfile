# Dockerfile for Render Deployment
FROM golang:1.26-bookworm AS builder

WORKDIR /app

# Install build dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    git \
    && rm -rf /var/lib/apt/lists/*

# Copy go mod files
COPY backend/go.mod backend/go.sum* ./

# Download dependencies
RUN go mod download

# Copy application source
COPY backend/*.go ./

# Build
RUN go build -o frontend_server .

# Runtime image
FROM alpine:latest

WORKDIR /app

# Copy binary
COPY --from=builder /app/frontend_server /app/frontend_server

# Set permissions
RUN chmod +x /app/frontend_server

# Expose port
EXPOSE 4242

# Run
CMD ["./frontend_server"]
