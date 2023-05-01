FROM golang:1.19

RUN apk update && apk add git

WORKDIR /code

COPY . .
RUN go env -w GOPROXY=https://goproxy.io,direct
RUN go env -w GO111MODULE=on
RUN go build -o *.go ./handler