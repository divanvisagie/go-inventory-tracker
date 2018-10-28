FROM flexconstructor/go-swagger




RUN mkdir -p /docs


ADD . /docs

# RUN go get -u github.com/go-swagger/go-swagger/cmd/swagger

WORKDIR /docs

EXPOSE 80

ENTRYPOINT swagger serve swagger.yml -F redoc --port=80 --host="0.0.0.0"