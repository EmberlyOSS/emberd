BINARY := emberd

.PHONY: all build install fmt vet tidy clean

all: build

build:
	go build -o $(BINARY) ./...

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
