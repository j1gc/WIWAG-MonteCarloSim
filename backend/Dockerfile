FROM golang:1.24 AS builder

WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/server ./cmd/

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin/server .

EXPOSE 4000

CMD ["/app/server"]