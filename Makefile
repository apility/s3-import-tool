BUILD_TIME := $(shell date +%s)
VERSION ?= development
ROOT_NAME ?= netflex-import

build_cli:
	go build \
		-o "dist/${ROOT_NAME}" \
		-ldflags "-X main.rootName=${ROOT_NAME} -X main.buildTime=${BUILD_TIME} -X main.versionNumber=${VERSION}" \
		./cmd/*.go