FROM golang:1.26.1-alpine AS builder

WORKDIR /app/backend

COPY backend/go.mod backend/go.sum ./

RUN go mod download

COPY backend/ ./

RUN go build -o /bin/api ./cmd/api

FROM alpine:3.21

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /bin/api /app/api

EXPOSE 8080

CMD ["/app/api"]
