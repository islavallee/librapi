FROM golang:1.15
WORKDIR /go/src/github.com/islavallee/librapi/
COPY  . .
RUN GOARCH=amd64 GOOS=linux go build

FROM alpine:3.12
MAINTAINER s.lavallee.pro@gmail.com
COPY --from=0 /go/src/github.com/islavallee/librapi/librapi .
EXPOSE 8080
USER nobody
ENTRYPOINT ["./librapi"]
