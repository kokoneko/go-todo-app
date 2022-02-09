FROM golang:1.17.6-alpine3.15

RUN apk add make gcc g++ git

RUN go get -u github.com/cosmtrek/air

RUN go env -w GO111MODULE=on

WORKDIR /api