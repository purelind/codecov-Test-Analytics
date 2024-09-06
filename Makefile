.PHONY: test coverage install-deps

GOPATH := $(shell go env GOPATH)
PATH := $(GOPATH)/bin:$(PATH)

test: install-deps
	go test -v -coverprofile=coverage.out ./... -count=10 | go-junit-report > junit.xml

coverage:
	go tool cover -html=coverage.out -o coverage.html

install-deps:
	go install github.com/jstemmer/go-junit-report@latest
