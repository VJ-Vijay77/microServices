FROM golang:alpine AS builder

RUN mkdir /app

WORKDIR /app

# COPY . /app/

RUN env GOOS=linux CGO_ENABLED=0 go build -o mongoBinary ./cmd/api/main.go

FROM alpine:3

WORKDIR /app

COPY --from=builder /app/mongoBinary .

CMD [ "/app/mongoBinary" ]