FROM golang:alpine

WORKDIR /go/src/github.com/edvakf/https-forward-proxy/client
COPY ./ /go/src/github.com/edvakf/https-forward-proxy/client

RUN go build -o app .

CMD ["/go/src/github.com/edvakf/https-forward-proxy/client/app"]
