# Go Inventory Tracker
An example service that does inventory tracking built with go-swagger


## Development Setup

### Generate from docs

```bash
brew tap go-swagger/go-swagger
brew install go-swagger
```

### Validating swagger docs

```bash
swagger validate ./swagger.yml
```

### Generating the server
swagger generate server -A inventory-tracker -f ./swagger.yml

## Running

Run server
```bash
make
```

Run swagger docs
```bash
make docs
```