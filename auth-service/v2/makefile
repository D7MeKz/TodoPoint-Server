GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOFMT=$(GOCMD) fmt
BINARY_NAME=app
SERVICE_VERSION=1.0.2

.PHONY: all build test clean

all: build

test:
	$(GOTEST) -v ./...

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

# Make linux binary
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v


clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

.PHONY: docker-build

docker-build:
	docker build -t auth:$(SERVICE_VERSION) -f Dockerfile ../..
