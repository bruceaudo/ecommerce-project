build:
	@go build -o bin/ecommerce-project cmd/main.go

run: build
	@./bin/ecommerce-project