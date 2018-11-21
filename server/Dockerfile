# Build stage
FROM divanvisagie/go-swagger-base
ARG SOURCE_ROOT=/go/src/github.com/divanvisagie/go-inventory-tracker
RUN mkdir -p /go/src/github.com/divanvisagie/go-inventory-tracker


ADD . ${SOURCE_ROOT}

WORKDIR ${SOURCE_ROOT}/server/cmd/inventory-tracker-server
RUN go get -v

WORKDIR ${SOURCE_ROOT}

RUN make build

EXPOSE 80

ENTRYPOINT ./goapp --port=80 --host="0.0.0.0"
