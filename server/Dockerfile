FROM golang:alpine

WORKDIR /go/src/github.com/edvakf/https-forward-proxy/server
COPY ./ /go/src/github.com/edvakf/https-forward-proxy/server

RUN go build -o app .

CMD ["/go/src/github.com/edvakf/https-forward-proxy/server/app"]
