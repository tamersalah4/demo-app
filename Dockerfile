# Dockerfile
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY main.go .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o server main.go

FROM alpine:3.20
COPY --from=builder /app/server /server
EXPOSE 8080
ENV VERSION="1.0.0"
CMD ["/server"]
