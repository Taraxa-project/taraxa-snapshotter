.DEFAULT_GOAL := build
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin
BUILD_DIR=build
EXECUTABLE=taraxa-snapshotter

help:
	@echo "This is a helper makefile for taraxa-snapshotter"
	@echo "Targets:"
	@echo "    lint:        run lint"
	@echo "    check:       run tests"
	@echo "    tidy         tidy go mod"
	@echo "    build        builds executable"

$(GOBIN)/golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOBIN) v1.54.2

.PHONY: tools
tools: $(GOBIN)/golangci-lint

lint: tools
	$(GOBIN)/golangci-lint run ./...

check:
	go test ./...

tidy:
	@echo "tidy..."
	go mod tidy

build: clean
	@mkdir -p $(BUILD_DIR)/linux_amd64
	env GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/linux_amd64/$(EXECUTABLE)

clean:
	@rm -rf $(BUILD_DIR)
	@mkdir $(BUILD_DIR)
	@touch $(BUILD_DIR)/.gitkeep