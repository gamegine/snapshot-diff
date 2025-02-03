FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.* .
RUN go mod download
COPY . .
RUN --mount=type=cache,id=go-build,target=/root/.cache/go-build go build -v cmd/api/api.go

# FROM scratch
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/api ./
ENTRYPOINT ["/app/api"]