FROM golang:alpine

WORKDIR /golang-rest-api

COPY . .

# to download all deps
RUN go mod download

# install swag for , openapi docs
RUN go install github.com/swaggo/swag/cmd/swag@latest

# swap docs gen
RUN swag init -g cmd/api/main.go -o docs

# app build process
RUN go build -o ./bin/api ./cmd/api

# db migration
RUN go build -o ./bin/migrate ./cmd/migrate

CMD ["/golang-rest-api/bin/api"]
EXPOSE 8080