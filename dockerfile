# dockerfile
FROM golang:1.22-alpine AS builder

COPY . /github.com/mchekalov/chat_server/source
WORKDIR /github.com/mchekalov/chat_server/source

RUN go mod download
RUN go build -o ./bin/chatserver cmd/server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/mchekalov/chat_server/source/bin/chatserver .
ADD .env .

CMD ["./chatserver"]