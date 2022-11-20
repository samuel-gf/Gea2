GO_FILES=$(wildcard *.go)

run:
	@go run $(GO_FILES)

fmt:
	@go fmt $(GO_FILES)

