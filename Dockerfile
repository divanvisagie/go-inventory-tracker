FROM golang

RUN mkdir -p /go/src/github.com/divanvisagie/go-inventory-tracker

ADD . /go/src/github.com/divanvisagie/go-inventory-tracker

WORKDIR /go/src/github.com/divanvisagie/go-inventory-tracker/cmd/inventory-tracker-server

RUN go get  -v 
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher --port=8080 --host="0.0.0.0" -run /go/src/github.com/divanvisagie/go-inventory-tracker/cmd/inventory-tracker-server -watch github.com/divanvisagie/go-inventory-tracker 

