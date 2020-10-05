
build:
	go build ./cmd/dip-cli
	go build ./cmd/dip-srv

lint:
	golangci-lint run

test:
	go test -v -cover ./...
