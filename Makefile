SHELL := /bin/bash

.PHONY:all
all: test

.PHONY:clean
clean:
	rm memprofile.out cpuprofile.out gaussian.test c.out || true

.PHONY:install-test-deps
install-test-deps:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin
	go install golang.org/x/vuln/cmd/govulncheck@latest

.PHONY:test
test: install-test-deps unit
	go vet ./...
	gofmt -e -l -d -s .
	go mod tidy
	golangci-lint run ./...
	govulncheck ./...

.PHONY:unit
unit:
	go test -coverprofile=c.out -covermode=atomic ./...

.PHONY:cover
cover: test
	go tool cover -func=c.out
	sed -i 's/github.com\/albertyw\/go-gaussian-elimination\///g' c.out

.PHONY:race
race:
	go test -race ./...

.PHONY:benchmark
benchmark:
	go test -bench=. -benchmem

.PHONY:cpuprof
cpuprof:
	go tool pprof -top cpuprofile.out | head -n 20

.PHONY:memprof
memprof:
	go tool pprof -top memprofile.out | head -n 20
