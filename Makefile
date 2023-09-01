GOPATH:=$(shell go env GOPATH)

.PHONY: build
build:
	go build -o task-app *.go

.PHONY: test
test:
	go test -v ./... -cover -race

.PHONY: vendor
vendor:
	go get ./...
	go mod vendor
	go mod verify

# ==============================================================================
# Swagger

swagger:
	@echo Starting swagger generating
	swag init -g *.go