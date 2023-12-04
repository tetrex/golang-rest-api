FROM golang:alpine

WORKDIR /myapp
COPY . .

RUN go build main.go

CMD ["/myapp/main"]
EXPOSE 8080