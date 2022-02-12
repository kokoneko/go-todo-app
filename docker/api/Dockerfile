FROM golang:1.17.6-alpine3.15

RUN apk add make gcc g++ git

RUN go get -u github.com/cosmtrek/air

RUN go get bitbucket.org/liamstask/goose/cmd/goose

ENV GO111MODULE on
ENV GOPATH /go

WORKDIR /go/src
COPY ./src .