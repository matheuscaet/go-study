# Go Study Project Makefile

# Variables
BINARY_NAME=go-study
MAIN_PATH=./cmd/main.go
GO_FILES=$(shell find . -name "*.go" -type f)

# Build the application
build:
	go build -o bin/$(BINARY_NAME) $(MAIN_PATH)

# Run the application
run:
	go run $(MAIN_PATH)

# Clean build artifacts
clean:
	go clean
	rm -f bin/$(BINARY_NAME)

# Create a new module (usage: make new-module name=modulename)
new-module:
	@if [ -z "$(number)" ] || [ -z "$(name)" ]; then \
		echo "Usage: make new-module number=<number> name=<module_name>"; \
		exit 1; \
	fi
	mkdir -p $(number)$(name)
	echo "package $(name)\n\nimport \"fmt\"\n\nfunc $(shell echo $(name) | sed 's/^./\U&/')() {\n\tfmt.Println(\"$(name) module\")\n}" > $(number)$(name)/$(name).go