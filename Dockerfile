FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY go.* .
RUN go mod download

COPY *.go .
RUN go build -o check

FROM alpine
WORKDIR /app
COPY --from=builder /app/check .