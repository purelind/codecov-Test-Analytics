.PHONY: test coverage

test:
	go test -v -coverprofile=coverage.out ./... -count=10 | go-junit-report > junit.xml

coverage:
	go tool cover -html=coverage.out -o coverage.html

install-deps:
	go get -u github.com/jstemmer/go-junit-report
