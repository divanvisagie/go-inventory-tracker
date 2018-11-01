main:
	docker-compose up
servedocs:
	swagger serve ./server/swagger.yml -F redoc --port=8081

generate:
	cd server; swagger generate server -A inventory-tracker -f ./swagger.yml

build:
	env GOOS=linux GOARCH=amd64 go build github.com/divanvisagie/go-inventory-tracker/server/cmd/inventory-tracker-server

dockerize:
	$(MAKE) build
	docker build -t divanvisagie/go-inventory-tracker:latest .
	$(MAKE) clean

clean:
	rm -rf load-testing/__pycache__
	rm inventory-tracker-server