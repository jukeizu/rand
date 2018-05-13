VERSION=$(shell git describe --tags)
BUILD=GOARCH=amd64 go build -v

.PHONY: all deps test build

all: deps test build

deps:
	go get -t -v ./...

test:
	go vet ./...
	go test -v -race ./...

build:
	for CMD in `ls cmd/listeners/commands`; do $(BUILD) -o bin/$$CMD-command-$(VERSION) ./cmd/listeners/commands/$$CMD; done
