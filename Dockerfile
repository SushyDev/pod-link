FROM golang:1.21.3-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build

CMD ["/app/pod-link"]
