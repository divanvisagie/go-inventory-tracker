FROM golang

RUN mkdir -p /go/src/github.com/divanvisagie/go-inventory-tracker

ADD . /go/src/github.com/divanvisagie/go-inventory-tracker

WORKDIR /go/src/github.com/divanvisagie/go-inventory-tracker/server/cmd/inventory-tracker-server

RUN go get  -v 
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

EXPOSE 80

ENTRYPOINT  watcher --port=80 --host="0.0.0.0" -run /go/src/github.com/divanvisagie/go-inventory-tracker/server/cmd/inventory-tracker-server -watch github.com/divanvisagie/go-inventory-tracker 

