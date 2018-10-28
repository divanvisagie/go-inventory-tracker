main:
	docker-compose up
servedocs:
	swagger serve ./server/swagger.yml -F redoc --port=8081

generate:
	cd server; swagger generate server -A inventory-tracker -f ./swagger.yml

clean:
	rm -rf load-testing/__pycache__