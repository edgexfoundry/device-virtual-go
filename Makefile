.PHONY: build test clean docker

GO=CGO_ENABLED=1 GO111MODULE=on go

# Don't need CGO_ENABLED=1 on Windows w/o ZMQ.
# If it is enabled something is invoking gcc and causing errors
ifeq ($(OS),Windows_NT)
  GO=CGO_ENABLED=0 GO111MODULE=on go
endif

MICROSERVICES=cmd/device-virtual

.PHONY: $(MICROSERVICES)

DOCKERS=docker_device_virtual_go
.PHONY: $(DOCKERS)

VERSION=$(shell cat ./VERSION 2>/dev/null || echo 0.0.0)
GIT_SHA=$(shell git rev-parse HEAD)
GOFLAGS=-ldflags "-X github.com/edgexfoundry/device-virtual-go.Version=$(VERSION)"

tidy:
	go mod tidy

build: $(MICROSERVICES)

cmd/device-virtual:
	$(GO) build $(GOFLAGS) -o $@ ./cmd

test:
	$(GO) test ./... -coverprofile=coverage.out
	$(GO) vet ./...
	gofmt -l $$(find . -type f -name '*.go'| grep -v "/vendor/")
	[ "`gofmt -l $$(find . -type f -name '*.go'| grep -v "/vendor/")`" = "" ]
	./bin/test-attribution-txt.sh

clean:
	rm -f $(MICROSERVICES)

docker: $(DOCKERS)

docker_device_virtual_go:
	docker build \
		--label "git_sha=$(GIT_SHA)" \
		-t edgexfoundry/device-virtual:$(GIT_SHA) \
		-t edgexfoundry/device-virtual:$(VERSION)-dev \
		.

vendor:
	CGO_ENABLED=0 GO111MODULE=on go mod vendor
