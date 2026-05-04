PROJECTNAME=gofilter
PWD_PROJECT=$(shell pwd)
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
BIN_DIR?=$(PWD_PROJECT)/bin

export GOBIN := $(BIN_DIR)

## help: List Make targets.
help: Makefile
	@echo " Choose a command run in $(PROJECTNAME):"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
.PHONY: help

## build: Compile the gofilter binary for the current OS/arch (output in bin/).
build:
	@echo "--> Building $(PROJECTNAME) binary for $(GOOS):$(GOARCH)"
	@mkdir -p "$(BIN_DIR)"
	@if [ "$(GOOS)" = "windows" ]; then \
		env CGO_ENABLED=1 go build -o "$(BIN_DIR)/$(PROJECTNAME).exe" ./cmd/gofilter; \
	else \
		env CGO_ENABLED=1 go build -o "$(BIN_DIR)/$(PROJECTNAME)" ./cmd/gofilter; \
	fi
	@echo "--> $(PROJECTNAME) for $(GOOS):$(GOARCH) built at $(BIN_DIR)"
.PHONY: build

## fixtures-demo: Write fixtures/demo.png for VHS / local demos.
fixtures-demo:
	go run ./scripts/generate_demo_png.go
.PHONY: fixtures-demo

## demo-gif: Build binary, regenerate demo PNG, run VHS (requires vhs, ttyd, ffmpeg on PATH).
demo-gif: build fixtures-demo
	@if [ "$(GOOS)" = "windows" ]; then \
		cp "$(BIN_DIR)/$(PROJECTNAME).exe" ./$(PROJECTNAME).exe; \
	else \
		cp "$(BIN_DIR)/$(PROJECTNAME)" ./$(PROJECTNAME); \
	fi
	vhs gofilter.tape
.PHONY: demo-gif
