# Build Stage
FROM golang:1.23-alpine AS builder

# Install certificates for HTTPS (Required for Prometheus API)
RUN apk update && apk add --no-cache ca-certificates git

WORKDIR /app

# Optimize caching by copying go.mod first
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build static binary
# -ldflags="-w -s": Remove debug info to reduce size
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/prometheus-mcp cmd/main.go

# Final Stage (Minimal Image)
FROM scratch

# Copy certificates from builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy binary
COPY --from=builder /app/prometheus-mcp /prometheus-mcp

# Expose nothing (MCP works over Stdio usually, but if HTTP transport is added later)
# EXPOSE 8080 

ENTRYPOINT ["/prometheus-mcp"]
