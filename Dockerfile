FROM golang:1.15
WORKDIR /go/src/github.com/islavallee/librapi/
COPY  . .
RUN GOARCH=amd64 GOOS=linux go build

FROM debian
MAINTAINER s.lavallee.pro@gmail.com

COPY --from=0 /go/src/github.com/islavallee/librapi/librapi .

RUN mkdir -p /var/librapi
RUN chown nobody:nogroup /var/librapi
USER nobody

EXPOSE 8080

CMD ["./librapi"]
