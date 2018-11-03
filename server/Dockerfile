# Build stage
FROM golang AS build-env
ARG SOURCE_ROOT=/go/src/github.com/divanvisagie/go-inventory-tracker
RUN mkdir -p /go/src/github.com/divanvisagie/go-inventory-tracker


ADD . ${SOURCE_ROOT}

WORKDIR ${SOURCE_ROOT}/server/cmd/inventory-tracker-server
RUN go get -v

WORKDIR ${SOURCE_ROOT}

RUN make build

RUN ls

# Final stage
FROM golang:alpine
ARG SOURCE_ROOT=/go/src/github.com/divanvisagie/go-inventory-tracker
COPY --from=build-env ${SOURCE_ROOT} /app/goapp

WORKDIR /app

EXPOSE 80

ENTRYPOINT ./goapp --port=80 --host="0.0.0.0"
