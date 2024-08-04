all: help

.PHONY: help
help: Makefile
	@echo
	@echo " Choose a make command to run"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo


# install packages needed
.PHONY: setup
setup:
	@go install github.com/air-verse/air@latest github.com/a-h/templ/cmd/templ@latest


.PHONY: build
build:
	@echo "Building..."
	@templ generate
	@tailwindcss -i static/css/input.css -o static/css/output.css
	@go build -o main cmd/main.go


# Run the application
.PHONY: run
run:
	@go run cmd/main.go


# Test the application
.PHONY: test
test:
	@echo "Testing..."
	@go test ./... -v


# Clean the binary
.PHONY: clean
clean:
	@echo "Cleaning..."
	@rm -f main


# Live Reload
.PHONY: watch
watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        make setup \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

