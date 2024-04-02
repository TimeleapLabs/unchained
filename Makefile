prod:
	GOOS=windows GOARCH=amd64 go build -o bin/unchained.windows.amd64.exe ./src/main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/unchained.darwin.amd64 ./src/main.go
	GOOS=linux GOARCH=amd64 go build -o bin/unchained.linux.amd64 ./src/main.go

	GOOS=windows GOARCH=arm64 go build -o bin/unchained.windows.arm64.exe ./src/main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/unchained.darwin.arm64 ./src/main.go
	GOOS=linux GOARCH=arm64 go build -o bin/unchained.linux.arm64 ./src/main.go

	find bin -type f -exec chmod u+x {} \;

strip:
	GOOS=windows GOARCH=amd64 go build -ldflags "-w -s" -o bin/unchained.windows.amd64.exe ./src/main.go
	GOOS=darwin GOARCH=amd64 go build -ldflags "-w -s" -o bin/unchained.darwin.amd64 ./src/main.go
	GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o bin/unchained.linux.amd64 ./src/main.go

	GOOS=windows GOARCH=arm64 go build -ldflags "-w -s" -o bin/unchained.windows.arm64.exe ./src/main.go
	GOOS=darwin GOARCH=arm64 go build -ldflags "-w -s" -o bin/unchained.darwin.arm64 ./src/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags "-w -s" -o bin/unchained.linux.arm64 ./src/main.go

	find bin -type f -exec chmod u+x {} \;

tools:
	go mod tidy
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
	  sh -s -- -b $(CURDIR)/bin v1.55.2
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

