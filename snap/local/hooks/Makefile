GO=go
HOOKS=cmd/configure/configure cmd/install/install
GOFLAGS=-ldflags="-s -w"

build: $(HOOKS)

cmd/configure/configure:
	$(GO) build $(GOFLAGS) -o $@ ./cmd/configure

cmd/install/install:
	$(GO) build $(GOFLAGS) -o $@ ./cmd/install

tidy:
	go mod tidy -compat=1.17

clean:
	rm -f $(HOOKS)

