.PHONY: test
test:
	@echo "Running tests..."
	@go test -v ./...

.PHONY: build
build:
	@echo "Building..."
	@go build -o app .

.PHONY: run
run:
	@echo "Running..."
	@go run main.go

