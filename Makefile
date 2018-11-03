main:
	docker-compose up
servedocs:
	swagger serve ./server/swagger.yml -F redoc --port=8081

generate:
	cd server; swagger generate server -A inventory-tracker -f ./swagger.yml

build:
	go build -o goapp github.com/divanvisagie/go-inventory-tracker/server/cmd/inventory-tracker-server

dockerize:
	docker build -t divanvisagie/go-inventory-tracker:latest  -f server/Dockerfile .
	docker build -t divanvisagie/inventory-tracker-migrations:latest -f migrations.Dockerfile .

clean:
	rm -rf load-testing/__pycache__
	rm goapp