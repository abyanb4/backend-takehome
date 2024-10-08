# Dockerfile
FROM golang:1.21.0-alpine

ENV TZ=UTC

# Install tzdata package to set timezone and curl for downloading
RUN apk update && apk add --no-cache tzdata curl tar

# Configure the timezone
RUN ln -sf /usr/share/zoneinfo/UTC /etc/localtime && echo "UTC" > /etc/timezone

# install git for golang-migrate
RUN apk add --no-cache git

# Determine architecture and download appropriate golang-migrate
RUN ARCH=$(uname -m) && \
    case $ARCH in \
        x86_64) \
            GOARCH=amd64 \
            ;; \
        aarch64|arm64) \
            GOARCH=arm64 \
            ;; \
        *) \
            echo "Unsupported architecture: $ARCH" \
            exit 1 \
            ;; \
    esac && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.1/migrate.linux-${GOARCH}.tar.gz | tar xz -C /usr/local/bin/ && \
    chmod +x /usr/local/bin/migrate

WORKDIR /app

# Copy the rest of the application source code
COPY . .

# remove and re-init package because then we can either use arm or x86
RUN rm go.mod
RUN rm go.sum
RUN go mod init backend-takehome
RUN go mod tidy

RUN chmod +x ./entrypoint.sh

# Build the Go application
RUN go build -o main ./cmd/http

# Use a non-root user for security
RUN adduser -D appuser
USER appuser

EXPOSE 8080

ENTRYPOINT ["/app/entrypoint.sh"]

