# Agents.md

This file provides guidance to LLM coding agents when working with code in this repository.

## Project Overview

A shopping list TUI application built with Go and [Bubble Tea](https://github.com/charmbracelet/bubbletea). The project follows Go module standards and uses the Bubble Tea framework for terminal-based user interfaces.

## Development Commands

### Build and Run

- `go run .` - Run the TUI application locally (current main.go implementation)
- `go build .` - Build the application binary
- `go run ./cmd/bubble-shopping-list` - Run the TUI application locally (future: main entry point will be moved here)
- `go build ./cmd/bubble-shopping-list` - Build the CLI binary (future)
- `go build -v ./...` - Build all packages with verbose output

### Testing and Quality

- `go test ./...` - Run all tests
- `go test ./... -race` - Run tests with race detection (recommended before pushing)
- `go test ./... -cover` - Run tests with coverage reporting
- `go vet ./...` - Run static analysis checks
- `go fmt ./...` - Format code according to Go standards

## Project Structure

The project follows Go module conventions with a planned structure:

- `cmd/bubble-shopping-list/` - CLI entry point (main application)
- `internal/` - Private application packages
  - `internal/app/` - Bubble Tea models and application logic
  - `internal/storage/` - Data persistence layer
  - `internal/ui/` - View helpers and UI components
  - `internal/testsupport/` - Testing utilities and fakes
- `assets/` - Static files, fixture data, example configurations

## Architecture Notes

- Uses Bubble Tea for TUI framework
- State management follows Bubble Tea's Model-View-Update pattern
- Testing emphasizes deterministic message handling with controlled seeds
- Module name: `github.com/theantichris/bubble-shopping-list`
- Go version: 1.25.1

## Code Style

- Use `gofmt` for consistent formatting (tabs, not spaces)
- Import grouping: stdlib, external, intra-module
- PascalCase for exported identifiers with descriptive comments
- lowerCamelCase for internal helpers
- Descriptive receivers like `model *Model`
- Lowercase filenames without underscores/dashes (e.g., `itemstore.go`)
