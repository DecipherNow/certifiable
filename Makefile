VERSIONS_PACKAGE := github.com/deciphernow/certifiable/versions

COMMIT := $(shell git rev-parse --verify --short HEAD 2> /dev/null || echo "UNKNOWN")
COMMIT_FLAG := -X $(VERSIONS_PACKAGE).commit=$(COMMIT)

VERSION := $(shell cat VERSION || echo "UNKNOWN")
VERSION_FLAG := -X $(VERSIONS_PACKAGE).version=$(VERSION)

.PHONY: build
build: vendor
	@echo "--> Building binary..."
	@CGO_ENABLED=0 go build -o bin/certifiable -ldflags "$(VERSION_FLAG) $(COMMIT_FLAG)" --mod=vendor

.PHONY: build.linux
build.linux: vendor
	@echo "--> Building binary..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/certifiable -ldflags "$(VERSION_FLAG) $(COMMIT_FLAG)" --mod=vendor

.PHONY: test
test: vendor
	@echo "--> Running tests..."
	@CGO_ENABLED=0 go test -v --mod=vendor ./...

.PHONY: vendor
vendor:
	@echo "--> Vendoring dependencies..."
	@CGO_ENABLED=0 go mod vendor

.PHONY: docker
docker: build.linux
	@echo "--> Building image..."
	@docker build -t deciphernow/certifiable:latest .
