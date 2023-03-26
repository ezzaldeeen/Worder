build:
	echo "Building..."
	go build -o bin/worder main.go

format:
	echo "Formatting..."
	go fmt

test:
	echo "Testing..."
	go test -v ./...

coverage:
	echo "Building Coverage Report..."
	go test --cover ./...

benchmark:
	echo "Benchmark..."
	go test --bench
