# Dependency Inversion Principle example

This app provide example how to use DIP and how it helps make well designed software.

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
