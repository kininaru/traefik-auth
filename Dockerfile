FROM golang:1.13-alpine as builder

RUN mkdir -p /go/src/github.com/Kininaru/traefik-auth
RUN apk add --no-cache git
ADD . /go/src/github.com/Kininaru/traefik-auth/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -installsuffix nocgo -o /traefik-auth github.com/Kininaru/traefik-auth/cmd

FROM scratch
COPY --from=builder /traefik-auth ./
ENTRYPOINT ["./traefik-auth"]