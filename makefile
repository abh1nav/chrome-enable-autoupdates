.PHONY: all build

APP = chrome-enable-autoupdates
VERSION = $(shell cat build/VERSION)

install:
	@GO111MODULE=on go get -u all
	@GO111MODULE=on go mod tidy
	@GO111MODULE=on go mod vendor

run:
	@go run -mod=vendor -ldflags "-X main.version=$(VERSION)" cmd/$(APP)/main.go

clean:
	@rm -Rf dist

build:
	@go build -mod=vendor -o dist/$(APP)-$(VERSION)-arm64 -ldflags "-X main.version=$(VERSION)" cmd/$(APP)/main.go

lint:
	@DOCKER_BUILDKIT=1 docker build -f build/Dockerfile.lint -t $(APP)-$(VERSION)-lint .
	@docker run -i --rm --name=$(APP)-$(VERSION)-lint $(APP)-$(VERSION)-lint