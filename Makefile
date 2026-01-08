GO := go1.25.1

build:
	$(GO) build ./cmd/...

install:
	$(GO) install ./cmd/...

install-debug:
	$(GO) install -ldflags "-X github.com/go/hook-scripts/internal/filefirewall.debugMode=true" ./cmd/...

clean:
	rm block-env

test:
	$(GO) test ./...
