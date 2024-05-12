build:
	@go build -o bin/turkey-address-api cmd/main.go

run: build
	@./bin/turkey-address-api

test:
	@go test -v ./..