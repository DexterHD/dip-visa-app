
build:
	go build ./cmd/dip-cli
	go build ./cmd/dip-srv

lint:
	golangci-lint run

test:
	go test -v -cover ./...

make ci-coverage-dependencies:
	go get github.com/axw/gocov/...
	go get github.com/AlekSi/gocov-xml

make ci-coverage-report: ci-coverage-dependencies
	go test -race -covermode=atomic -coverprofile=coverage.txt ./...
	gocov convert coverage.txt | gocov-xml > coverage.xml
