# Build stage
FROM golang:1.24-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git make

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build argument for version
ARG VERSION=dev

# Build the binary
RUN go build -ldflags="-X main.Version=${VERSION}" -o ezpw cmd/ezpw/main.go

# Runtime stage
FROM alpine:latest

# Install runtime dependencies
RUN apk add --no-cache \
    ca-certificates \
    chromium \
    chromium-chromedriver \
    nodejs \
    npm

# Install Playwright
RUN npm install -g playwright && \
    playwright install chromium

# Create non-root user
RUN addgroup -g 1000 ezpw && \
    adduser -u 1000 -G ezpw -D ezpw

# Copy binary from builder
COPY --from=builder /app/ezpw /usr/local/bin/ezpw

# Set working directory
WORKDIR /workspace

# Change to non-root user
USER ezpw

# Default command
ENTRYPOINT ["ezpw"]
CMD ["--help"]