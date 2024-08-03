all: help

.PHONY: help
help: Makefile
	@echo
	@echo " Choose a make command to run"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

## setup: install packages needed
.PHONY: setup
setup:
	go install github.com/air-verse/air@latest github.com/a-h/templ/cmd/templ@latest

## build: build a binary
.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./run-app ./cmd/main.go

## docker-build: build project into a docker container image
.PHONY: docker-build
docker-build:
	GOPROXY=direct docker buildx build -t ${name} .

## docker-run: run project in a container
.PHONY: docker-run
docker-run:
	docker run -it --rm -p 8080:8080 ${name}

## start: build and run local project
.PHONY: start
start:
	air
	
## css: build tailwindcss
.PHONY: css
css:
	tailwindcss -i static/input.css -o static/output.css --minify

## css-watch: watch build tailwindcss
.PHONY: css-watch
css-watch:
	tailwindcss -i static/input.css -o static/output.css --watch
