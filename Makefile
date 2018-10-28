main:
	docker-compose up
servedocs:
	swagger serve ./server/swagger.yml -F redoc --port=8081

clean:
	rm -rf load-testing/__pycache__