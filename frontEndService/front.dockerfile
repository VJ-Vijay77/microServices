FROM golang:alpine AS builder

RUN mkdir /app

WORKDIR /app

COPY frontEnd /app

FROM alpine:3

WORKDIR /app

COPY --from=builder /app/frontEnd .

CMD [ "/app/frontEnd" ]