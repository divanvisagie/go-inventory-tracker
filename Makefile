main:
	go install ./cmd/inventory-tracker-server
	inventory-tracker-server --port=8080
docs:
	swagger serve swagger.yml -F redoc