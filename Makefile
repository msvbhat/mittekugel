GO_FILES := $(shell find . -type f -name '*.go' | grep -v /vendor/ | grep -v _test.go)
PKG_LIST := $(shell go list ./... | grep -v /vendor/)
BINARY := mittekugel

.PHONY: all dep build clean test coverage lint

all: build

dep: ## Get the dependencies
	@go get -v -d ./...

lint: ## Run the golint
	@golint -set_exit_status ${PKG_LIST}

test: ## Run unit tests
	@go test -short ${PKG_LIST}

build: ## Build the static binary
	@env CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -i -o ${BINARY}

clean: ## Remove the previous build
	@rm -f ${BINARY}
