# syntax=docker/dockerfile:1

FROM golang:1.20-alpine
WORKDIR /app

COPY . .

RUN go build -o script .

CMD ["./script"]
