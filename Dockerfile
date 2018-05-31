FROM golang:1.10.2-alpine3.7

WORKDIR /go/src/github.com/seregayoga/shorturl

COPY . .

RUN go install github.com/seregayoga/shorturl/cmd/shorturl

CMD ["shorturl"]
