clean:
	rm -rf agent

build: clean
	@cd exec && go build -o ../agent agent.go
	@sudo mv agent /usr/local/bin

run: build
	./agent

test: 
	@cd exec && go test ./src/...
