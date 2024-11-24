build:
	mkdir -p bin
	go build -o ./bin ./cmd/receipt-processor

run: build
	./bin/receipt-processor
