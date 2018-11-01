FROM alpine

RUN mkdir -p /app
COPY inventory-tracker-server /app

WORKDIR /app

EXPOSE 80

ENTRYPOINT ./inventory-tracker-server --port=80 --host="0.0.0.0"
