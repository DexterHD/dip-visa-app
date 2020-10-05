
build:
	go build ./cmd/dip-cli
	go build ./cmd/dip-srv

lint:
	golangci-lint run

test:
	go test -v -cover ./...

ci-coverage-dependencies:
	go get github.com/axw/gocov/...
	go get github.com/AlekSi/gocov-xml
	go mod tidy

ci-coverage-report: ci-coverage-dependencies
	go test -race -covermode=atomic -coverprofile=coverage.txt ./...
	gocov convert coverage.txt | gocov-xml > coverage.xml

clean:
	rm ./coverage.txt
	rm ./coverage.xml
	rm ./dip-srv
	rm ./dip-cli
