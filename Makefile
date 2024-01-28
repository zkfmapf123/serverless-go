build:
	@cd exec && go build -o ../main main.go

run: build
	./main

test:
	@cd exec && go test ./...

test-w:
	@cd exec && gow test -v