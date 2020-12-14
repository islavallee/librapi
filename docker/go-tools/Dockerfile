FROM golang:1.15

MAINTAINER s.lavallee.pro@gmail.com

ENV GO111MODULE=on

# https://github.com/cespare/reflex run a command when files change
RUN go get \
 github.com/cespare/reflex \
 golang.org/x/lint/golint \
 github.com/fzipp/gocyclo/cmd/gocyclo \
 github.com/securego/gosec/v2/cmd/gosec

RUN mkdir -p /var/librapi
WORKDIR /go/src/github.com/islavallee/librapi/

EXPOSE 8080

