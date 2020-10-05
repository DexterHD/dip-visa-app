# Dependency Inversion Principle Example Application

![golangci-lint](https://github.com/idexter/dip-visa-app/workflows/golangci-lint/badge.svg)

This application provides examples of DIP and how it helps to make software design better.

### Examples

```bash
go run ./cmd/dip-cli --id 1
go run ./cmd/dip-cli --id 0
```

```bash
go run ./cmd/dip-srv
curl -v -X GET "http://localhost:8080/?id=1"
curl -v -X GET "http://localhost:8080/?id=3"
```
