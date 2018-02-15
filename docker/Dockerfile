FROM golang:alpine
MAINTAINER Tim Fall <tim@vapor.io>

RUN apk update && apk add git

RUN go get -d -v github.com/vapor-ware/synse-cli

WORKDIR /go/src/github.com/vapor-ware/synse-cli

RUN go get -v github.com/golang/dep/cmd/dep
RUN dep ensure -v --vendor-only

RUN cd cmd/synse/ && go build -v -o /go/bin/synse

ENTRYPOINT ["synse"]
CMD ["--help"]
