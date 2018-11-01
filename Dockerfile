FROM alpine

RUN mkdir -p /app
COPY app /app

WORKDIR /app

EXPOSE 80

ENTRYPOINT ./app --port=80 --host="0.0.0.0"
