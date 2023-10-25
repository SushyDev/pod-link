FROM golang:1.21.3-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o main .

CMD ["/app/main"]
