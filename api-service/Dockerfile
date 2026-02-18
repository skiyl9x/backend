# Build Stage
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o api-service api-service.go

# Run Stage
FROM alpine:latest
WORKDIR /app/
COPY --from=builder /app/api-service .
EXPOSE 8080
ENTRYPOINT ["./api-service"]
