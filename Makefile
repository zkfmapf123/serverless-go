build:
	@go build -o main ./bin/main.go

dev:
	@go run ./bin/main.go
	
run:
	./main