FROM golang:latest


ADD . /go/src/github.com/Gujarats/API-Golang

RUN go get github.com/Gujarats/API-Golang
RUN go install github.com/Gujarats/API-Golang


ENTRYPOINT /go/bin/API-Golang

EXPOSE 8080
