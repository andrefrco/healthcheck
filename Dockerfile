FROM golang:alpine

ENV APPNAME healthcheck

RUN apk add git
RUN mkdir -p /go/src /go/bin /go/github.com/$APPNAME && chmod -R 777 /go

ENV GOPATH /go
ENV PATH /go/bin:$PATH

WORKDIR /go/github.com/$APPNAME

ADD . /go/github.com/$APPNAME

RUN go get github.com/gorilla/mux

CMD ["go", "run", "main.go"]
