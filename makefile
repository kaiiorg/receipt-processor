build:
	mkdir -p bin
	go mod tidy
	go build -o ./bin ./cmd/receipt-processor

run: build
	./bin/receipt-processor

test:
	go test ./...
