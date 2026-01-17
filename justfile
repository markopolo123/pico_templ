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

# Run pre-commit on all files
pre-commit:
    prek run --all-files

# Install pre-commit hooks
pre-commit-install:
    prek install

# Check release configuration (dry run)
release-check:
    goreleaser check

# Build snapshot release locally (no publish)
release-snapshot:
    goreleaser release --snapshot --clean

# Create and push a new release tag
release version:
    @echo "Creating release {{version}}..."
    git tag -a "v{{version}}" -m "Release v{{version}}"
    git push origin "v{{version}}"
    @echo "Release v{{version}} tagged and pushed. GitHub Actions will build and publish."

# Build release locally with 1Password GitHub token
release-local:
    GITHUB_TOKEN=$(op read "op://homelab/pico GitHub Personal Access Token/token") goreleaser release --clean
