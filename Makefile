
build:
	go build ./cmd/dip-cli
	go build ./cmd/dip-srv

test:
	go test -v -cover ./...
