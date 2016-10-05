FROM golang:alpine

WORKDIR /go/src/github.com/edvakf/nginx-stream-forward-proxy-example/client
COPY ./ /go/src/github.com/edvakf/nginx-stream-forward-proxy-example/client

RUN go build -o app .

CMD ["/go/src/github.com/edvakf/nginx-stream-forward-proxy-example/client/app"]
