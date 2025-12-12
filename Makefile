BINARY := emberd

.PHONY: all build install fmt vet tidy clean

all: build

build:
	# Build the main package at module root. Using `./...` with `-o <file>`
	# attempts to write multiple packages to a single file which fails.
	go build -v -o $(BINARY) .

install: build
	install -m 0755 $(BINARY) /usr/local/bin/$(BINARY)

fmt:
	gofmt -s -w .

vet:
	go vet ./...

tidy:
	go mod tidy

clean:
	rm -f $(BINARY)
