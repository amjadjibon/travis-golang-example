BINARY_NAME=travis-golang-example

run:
	go run main.go

test:
	go test ./... -v

cover:
	go test ./... -coverprofile=cover.out -v

cover-html:
	go tool cover -html=cover.out

build:
	go build -o bin/$(BINARY_NAME) main.go

build-docker:
	docker build -t "$(DOCKER_USERNAME)/$(BINARY_NAME)" .

run-docker:
	docker run -p 8080:8080 $(BINARY_NAME)

clean:
	go clean
	rm -f bin/$(BINARY_NAME)

docker-clean:
	docker rmi "$(DOCKER_USERNAME)$(BINARY_NAME)"

docker-push:
	echo "$(DOCKER_PASSWORD)" | docker login -u "$(DOCKER_USERNAME)" --password-stdin
	docker push "$(DOCKER_USERNAME)$(BINARY_NAME)"
