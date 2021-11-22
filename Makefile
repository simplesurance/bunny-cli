BIN		:= bunny-cli
VERSION		:= $(shell git describe --tags --always --dirty)
LDFLAGS	 	:= "-X github.com/simplesurance/bunny-cli/internal/command.version=$(VERSION)"
BUILDFLAGS 	:= -trimpath -ldflags=$(LDFLAGS)

default: build

.PHONY: build
build:
	$(info * compiling)
	go build $(BUILDFLAGS) -o $(BIN) main.go

.PHONY: check
check:
	$(info * running golangci-lint code checks)
	golangci-lint run
