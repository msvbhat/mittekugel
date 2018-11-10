GO_FILES := $(shell find . -type f -name '*.go' | grep -v /vendor/ | grep -v _test.go)
PKG_LIST := $(shell go list ./... | grep -v /vendor/)
BINARY := mittekugel

.PHONY: all dep build clean test coverage lint

all: build

dep: ## Get the dependencies
	@go get -v -d ./...

lint: ## Run the golint
	@golint -set_exit_status ${PKG_LIST}

build: ## Build the binary file
	@env GOOS=linux go build -i -o ${BINARY}

clean: ## Remove the previous build
	@rm -f ${BINARY}
