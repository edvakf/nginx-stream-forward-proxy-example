FROM golang:alpine

WORKDIR /go/src/github.com/edvakf/nginx-stream-forward-proxy-example/server
COPY ./ /go/src/github.com/edvakf/nginx-stream-forward-proxy-example/server

RUN go build -o app .

CMD ["/go/src/github.com/edvakf/nginx-stream-forward-proxy-example/server/app"]
