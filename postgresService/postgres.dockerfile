FROM golang:alpine AS builder

RUN mkdir /app

WORKDIR /app

COPY . /app/
COPY go.mod .
COPY go.sum .
RUN go mod download

RUN env GOOS=linux CGO_ENABLED=0 go build -o postgresBinary ./cmd/api/

FROM alpine:3

WORKDIR /app

COPY --from=builder /app/postgresBinary .

CMD [ "/app/postgresBinary" ]