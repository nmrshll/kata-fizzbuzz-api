# FizzBuzz API in Go

## Quick start

Run `make dev` to start server with docker (requires: docker, docker-compose) 

(alternative: run `go run cmd/api/main.go` (requires go 1.11+))

Try a request to locahost:8080:

`http POST localhost:8080/fizzbuzz limit:=15 int1:=3 int2:=5 str1="fizz" str2="buzz"`(requires httpie)

#### Requirements

- Go 1.11+
- docker
- docker-compose

## Testing

Run tests with `make test`

## License: MIT
