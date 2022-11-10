build:
	@go build -o ./bin/pingo ./cmd/...

run: build
	@./bin/pingo

tests:
	@go test -v ./test/...