GO := go1.25.1

build:
	$(GO) build ./cmd/...

install:
	$(GO) install ./cmd/...
