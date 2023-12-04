FROM golang:alpine

WORKDIR /golang-rest-api
COPY . .

RUN go build -o ./bin/api ./cmd/api

CMD ["/golang-rest-api/bin/api"]
EXPOSE 8080