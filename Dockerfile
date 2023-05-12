FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go build ./...

# FROM scratch
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/snapshot-diff ./