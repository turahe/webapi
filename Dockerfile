# Build Stage
FROM golang:1.24.2-alpine AS builder
LABEL maintainer="Nur Wachid <wachid@outlook.com>"
ENV CGO_ENABLED=1
ENV GO111MODULE=on

RUN apk add --no-cache git gcc g++

WORKDIR /app

# Copy Go Modules files
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the rest of the application source code
COPY . .

RUN go build -o /app/webapi /app/main.go
RUN ls /app -lah
# Run stage
FROM alpine:latest
RUN apk add --no-cache ca-certificates

ENV TZ=Asia/Jakarta
WORKDIR /app

# Copy the built binary
COPY --from=builder /app/webapi /app/webapi
COPY --from=builder /app/config/config.example.yaml /app/config/config.yaml
# Ensure executable permissions
RUN chmod +x /app/webapi

EXPOSE 8000

ENTRYPOINT ["/app/webapi", "server"]
