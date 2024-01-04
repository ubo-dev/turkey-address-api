build:
	@go build -o bin/turkey-address-api

run: build
	@./bin/turkey-address-api

test:
	@go test -v ./..