FROM golang:alpine

COPY . /go/app

WORKDIR /go/app

RUN go mod download
RUN go build -o ./.bin/server ./cmd/api/main.go

CMD ["./.bin/server"]