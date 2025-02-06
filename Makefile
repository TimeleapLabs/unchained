# Variables
BIN_DIR := bin
CMD := ./cmd/main.go
ARCHS := amd64 arm64
OS_LIST := windows darwin linux

# Targets
all: prod

prod: build

strip: LD_FLAGS=-ldflags "-w -s"
strip: build

darwin: OS=darwin
darwin: build

linux: OS=linux
linux: build

windows: OS=windows
windows: build

# Unified build rule
build:
	@for os in $(if $(OS),$(OS),$(OS_LIST)); do \
		for arch in $(ARCHS); do \
			output="$(BIN_DIR)/timeleap.$$os.$$arch"; \
			[ $$os = "windows" ] && output+=".exe"; \
			echo "Building $$output $(if $(LD_FLAGS),with flags: $(LD_FLAGS),)"; \
			GOOS=$$os GOARCH=$$arch go build $(LD_FLAGS) -o $$output $(CMD); \
		done; \
	done

exec:
	@echo "Setting executable permissions for files in $(BIN_DIR)"
	@find $(BIN_DIR) -type f -exec chmod u+x {} \;

tools:
	go mod tidy
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
 	  sh -s -- -b $(CURDIR)/bin v1.61.0

	go get golang.org/x/tools/cmd/goimports
	go install github.com/kisielk/errcheck@latest

check: errors imports fmt lint

errors:
	errcheck -ignoretests -ignoregenerated -blank ./...

lint:
	$(CURDIR)/bin/golangci-lint run

imports:
	goimports -l -w .

fmt:
	go fmt ./...

# Prevent Make from misinterpreting `build-specific` arguments
.PHONY: all prod strip darwin linux windows exec build
