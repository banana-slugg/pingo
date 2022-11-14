build:
	@go build -o ./bin/pingo ./cmd/...

run: build
	@./bin/pingo

tests:
	@go test -v ./test/...

windows:
	@GOOS=windows GOARCH=amd64 go build -o ./bin/pingo.exe ./cmd/...