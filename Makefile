.PHONY: build test unittest lint clean docker

GO=CGO_ENABLED=1 GO111MODULE=on go

# Don't need CGO_ENABLED=1 on Windows w/o ZMQ.
# If it is enabled something is invoking gcc and causing errors
ifeq ($(OS),Windows_NT)
  GO=CGO_ENABLED=0 GO111MODULE=on go
endif

# see https://shibumi.dev/posts/hardening-executables
CGO_CPPFLAGS="-D_FORTIFY_SOURCE=2"
CGO_CFLAGS="-O2 -pipe -fno-plt"
CGO_CXXFLAGS="-O2 -pipe -fno-plt"
CGO_LDFLAGS="-Wl,-O1,–sort-common,–as-needed,-z,relro,-z,now"

MICROSERVICES=cmd/device-virtual

.PHONY: $(MICROSERVICES)

ARCH=$(shell uname -m)

DOCKERS=docker_device_virtual_go
.PHONY: $(DOCKERS)

VERSION=$(shell cat ./VERSION 2>/dev/null || echo 0.0.0)
GIT_SHA=$(shell git rev-parse HEAD)

ifeq ($(OS),Windows_NT)
	GOFLAGS=-ldflags "-X github.com/edgexfoundry/device-virtual-go.Version=$(VERSION)"
	CGOFLAGS=-ldflags "-X github.com/edgexfoundry/device-virtual-go.Version=$(VERSION)"
else
	GOFLAGS=-ldflags "-X github.com/edgexfoundry/device-virtual-go.Version=$(VERSION)" -trimpath -mod=readonly
	CGOFLAGS=-ldflags "-linkmode=external -X github.com/edgexfoundry/device-virtual-go.Version=$(VERSION)" -trimpath -mod=readonly -buildmode=pie
endif

build: $(MICROSERVICES)

tidy:
	go mod tidy

cmd/device-virtual:
	$(GO) build $(CGOFLAGS) -o $@ ./cmd


unittest:
	$(GO) test ./... -coverprofile=coverage.out

lint:
	@which golangci-lint >/dev/null || echo "WARNING: go linter not installed. To install, run\n  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b \$$(go env GOPATH)/bin v1.42.1"
	@if [ "z${ARCH}" = "zx86_64" ] && which golangci-lint >/dev/null ; then golangci-lint run --config .golangci.yml ; else echo "WARNING: Linting skipped (not on x86_64 or linter not installed)"; fi

test: unittest lint
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
