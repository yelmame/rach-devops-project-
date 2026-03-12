#!/bin/sh
set -e

echo "=== EntryPoint Started ==="
echo "Go version: $(go version)"
echo "solc version: $(solc --version)"

# Build the app
if [ -f "main.go" ]; then
    echo "Building Go app..."
    go mod download
    go build -o /app/go-sdk-app main.go
    echo "Build complete: /app/go-sdk-app"
fi

# Start the app
if [ -f "/app/go-sdk-app" ]; then
    echo "Starting Go SDK App on :8080..."
    exec /app/go-sdk-app
else
    echo "No app found, starting shell"
    exec sh
fi
