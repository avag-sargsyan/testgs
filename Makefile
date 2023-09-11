GO_BUILD_CMD=go build
BINARY_NAME=bin/app
MAIN_FILE=cmd/app/main.go
DOCKER_IMG=test-gs

.PHONY: build
build:
	$(GO_BUILD_CMD) -o $(BINARY_NAME) $(MAIN_FILE)

.PHONY: run
run:
	go run $(MAIN_FILE)

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	rm -f $(BINARY_NAME)

.PHONY: docker-build
docker-build:
	docker build -t $(DOCKER_IMG) .

.PHONY: docker-run
docker-run:
	docker run -p 8080:8080 $(DOCKER_IMG)

.PHONY: docker-up
docker-up: docker-build docker-run
