GO_FILES=$(wildcard *.go)
TGT=bin/server bin/client


all: bin $(TGT)

test: test.go
	go run test.go board.go cell.go

bin:
	@mkdir -p bin

bin/server: server.go board.go cell.go command.go test.go
	go build -o $@ $^

bin/client: client.go board.go cell.go command.go test.go token.go
	go build -o $@ $^

fmt:
	go fmt $(GO_FILES)

clean:
	rm -Rf $(TGT)

open_sc:
	nvim -O server.go client.go
