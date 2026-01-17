# https://just.systems

set shell := ["flox", "activate", "--", "bash", "-c"]

default:
    @just --list

# Generate templ files
generate:
    templ generate

# Run all tests
test: generate
    go test ./...

# Run tests with verbose output
test-v: generate
    go test -v ./...

# Build the project
build: generate
    go build ./...

# Format all files
fmt:
    templ fmt .
    go fmt ./...

# Run linter
lint:
    golangci-lint run

# Clean generated files
clean:
    find . -name '*_templ.go' -delete

# Run tests for a specific package
test-pkg pkg: generate
    go test -v ./{{pkg}}/...

# Generate and watch for changes
watch:
    templ generate --watch
