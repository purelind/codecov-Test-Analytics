.PHONY: test

test:
	go test -v ./... -count=10 | go-junit-report > junit.xml

install-deps:
	go get -u github.com/jstemmer/go-junit-report
