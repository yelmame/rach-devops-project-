# syntax=docker/dockerfile:1

# Use Go 1.23.6 (Debian 12) image
FROM golang:1.23.6-bookworm

# Install dependencies
RUN apt-get update && \
    apt-get install -y curl wget unzip git && \
    rm -rf /var/lib/apt/lists/*

# Install solc 0.8.20
ENV SOLC_VERSION=0.8.20
RUN curl -fsSL -o solc-amd64 "https://github.com/ethereum/solidity/releases/download/v${SOLC_VERSION}/solc-static-linux" && \
    chmod +x solc-amd64 && \
    mv solc-amd64 /usr/local/bin/solc

# Working directory
WORKDIR /root/go/src/app

# Copy ALL source files (your app + entrypoint)
COPY main.go go.mod entrypoint.sh ./
COPY k8s ./k8s

RUN chmod +x entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
CMD ["sh"]
