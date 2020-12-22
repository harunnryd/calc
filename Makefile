.PHONY: default help install build build-windows run run-windows

APP_NAME     = calc
VERSION      ?= $(shell git describe --always --tags)
GIT_COMMIT   ?= $(shell git rev-parse --short HEAD)
GIT_BRANCH   ?= $(shell git rev-parse --abbrev-ref HEAD)
BUILD_DATE   = $(shell date '+%Y-%m-%d-%H:%M:%S')

default: help

help:
	@echo 'These are common ${APP_NAME} commands used in various situations:'
	@echo
	@echo 'Usage:'
	@echo '   make install             Install all project dependencies.'
	@echo '   make build               Build the project.'
	@echo '   make build-windows       Build the project for Windows.'
	@echo '   make run ARGS=           Running the application with supplied arguments.'
	@echo '   make run-windows ARGS=   Running the application supplied arguments on Windows.'
	@echo '   make clean               Clean the directory tree.'

install:
	@echo "Install project dependencies"
	go get -v ./...

build: install
	@echo "Building ${APP_NAME} $(VERSION) $(GIT_COMMIT)"
	go build -ldflags "-X github.com/harunnryd/calc/version.Number=${VERSION} -X github.com/harunnryd/calc/version.GitCommit=${GIT_COMMIT} -X github.com/harunnryd/calc/version.BuildDate=${BUILD_DATE}" -o bin/${APP_NAME}

build-windows: install
	@echo "Building ${APP_NAME}.exe $(VERSION) $(GIT_COMMIT)"
	go build -ldflags "-X github.com/harunnryd/calc/version.Number=${VERSION} -X github.com/harunnryd/calc/version.GitCommit=${GIT_COMMIT} -X github.com/harunnryd/calc/version.BuildDate=${BUILD_DATE}" -o bin/${APP_NAME}.exe

clean:
	@echo "Removing ${APP_NAME} ${VERSION}"
	@test ! -e bin/${APP_NAME} || rm bin/${APP_NAME}