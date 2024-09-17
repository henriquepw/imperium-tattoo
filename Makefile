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
	@tailwindcss -i static/css/input.css -o static/css/output.css --minify
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



# run templ generation in watch mode to detect all .templ files and 
# re-create _templ.txt files on change, then send reload event to browser. 
# Default url: http://localhost:7331
watch/templ:
	templ generate --watch --proxy="http://localhost:3000" --open-browser=false


# run air to detect any go file changes to re-build and re-run the server.
watch/server:
	air \
	--build.cmd "go build -o .tmp/main main.go" \
	--build.bin ".tmp/main" \
	--build.delay "100" \
	--build.exclude_dir ".tmp" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true


# run tailwindcss to generate the styles.css bundle in watch mode.
watch/tailwind:
	tailwindcss -i static/css/input.css -o static/css/output.css --minify --watch


# watch for any js or css change in the assets/ folder, then reload the browser via templ proxy.
watch/sync_assets:
	air \
	--build.cmd "templ generate --notify-proxy" \
	--build.bin "true" \
	--build.delay "100" \
	--build.exclude_dir "" \
	--build.exclude_regex "output.css" \
	--build.include_dir "static" \
	--build.include_ext "js,css"


# start all 4 watch processes in parallel.
.PHONY: watch
watch: 
	make -j4 watch/templ watch/server watch/tailwind watch/sync_assets

