build:
	@cd exec && go build -o ../agent agent.go

run: build
	./agent

test: 
	@cd exec && go test ./src/... -v
