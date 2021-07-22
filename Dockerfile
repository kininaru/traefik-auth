FROM golang:1.16-alpine as builder

RUN mkdir -p /go/src/github.com/Kininaru/traefik-auth
RUN apk add --no-cache git
ADD . /go/src/github.com/Kininaru/traefik-auth/
WORKDIR /go/src/github.com/Kininaru/traefik-auth/
RUN GO111MODULE=on GOPROXY=https://goproxy.cn go build
ENTRYPOINT ["./traefik-auth"]