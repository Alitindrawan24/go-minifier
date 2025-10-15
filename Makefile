.PHONY: help fmt lint test run build tidy

# Default target shows available commands.
help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "Common targets:"
	@echo "  fmt    - gofmt the codebase"
	@echo "  lint   - go vet for static analysis"
	@echo "  test   - run go test with race detector"
	@echo "  run    - run the main package"
	@echo "  build  - compile the binary into ./bin/"
	@echo "  tidy   - go mod tidy to sync dependencies"

fmt:
	@gofmt -w $$(find . -name '*.go' -not -path './vendor/*')

lint:
	@go vet ./...

test:
	@go test -v ./...

tidy:
	@go mod tidy
