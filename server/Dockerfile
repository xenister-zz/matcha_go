FROM golang:latest

RUN mkdir -p /go/src/matcha_go
RUN go get github.com/johnnadratowski/golang-neo4j-bolt-driver
RUN go get github.com/oxequa/realize
RUN go get github.com/xenister/matcha_go
WORKDIR /go/src/github.com/xenister/matcha_go
RUN go get
EXPOSE 81

WORKDIR /go/src/github.com/xenister/matcha_go
CMD ["realize", "start"]