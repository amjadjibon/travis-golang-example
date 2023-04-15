BINARY_NAME=travis-golang-example

test:
	go test ./... -v

cover:
	go test ./... -coverprofile=cover.out -v

cover-html:
	go tool cover -html=cover.out

build:
	go build -o bin/$(BINARY_NAME) main.go
