BINARY_NAME := puzzle

deps:
	go mod vendor

build:	deps
	go build -o ./dist/$(BINARY_NAME) ./cmd/puzzle-challenge

test:	deps
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	rm coverage.out

run:	build
	./dist/$(BINARY_NAME)