FROM golang:latest

RUN mkdir -p /go/src/matcha_go
RUN go get github.com/johnnadratowski/golang-neo4j-bolt-driver
RUN go get github.com/oxequa/realize
WORKDIR /go/src/matcha_go
RUN go get
EXPOSE 81

WORKDIR /go/src/matcha_go
CMD ["realize", "start"]