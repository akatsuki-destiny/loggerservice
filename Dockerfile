FROM golang:1.20.5-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o loggerservice ./cmd/main.go

RUN chmod +x /app/loggerservice

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/loggerservice /app
COPY ./dev.env ./dev.env

EXPOSE 3003:3003

CMD ["/app/loggerservice"]