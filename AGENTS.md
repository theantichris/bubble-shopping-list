# Repository Guidelines

## Project Structure & Module Organization

Source code lives under `cmd/bubble-shopping-list` for the CLI entry point and `internal/` for reusable packages. Use subpackages like `internal/app` for Bubble Tea models, `internal/storage` for persistence, and `internal/ui` for view helpers. Keep assets such as fixture data or example configuration in `assets/`. Tests sit next to the code they exercise as `_test.go` files; integration flows can live in `internal/testsupport`.

## Build, Test, and Development Commands

- `go run ./cmd/bubble-shopping-list` launches the TUI locally with live reload of your changes.
- `go build ./cmd/bubble-shopping-list` produces the CLI binary under `./cmd/bubble-shopping-list`.
- `go test ./...` runs the full suite; add `-race` before pushing for data races.
- `go vet ./...` performs static checks aligned with Bubble Tea idioms.
  Add new tools to `tools.go` if they must be tracked in modules.

## Coding Style & Naming Conventions

Format code with `gofmt` (run `go fmt ./...`) before committing; tabs are the default. Group imports as stdlib, external, intra-module. Exported identifiers use PascalCase and include descriptive comments; internal helpers stay lowerCamelCase. Use descriptive receivers such as `model *Model` for components. Keep filenames lowercase with no underscores or dashes for multiword segments, e.g., `itemstore.go`.

## Testing Guidelines

Unit tests rely on Go’s `testing` package and should mirror package names. Name tests `TestFeatureBehavior` and subtests using t.Run(). Subtests that are independent should start with t.Parallel(). Prefer fakes in `internal/testsupport`. Capture Bubble Tea message handling with deterministic seeds. Run `go test ./... -cover` and maintain coverage on state transition logic; add regression tests when fixing bugs.

## Commit & Pull Request Guidelines

Write imperative, present-tense subject lines under 72 characters, e.g., `add optimistic item removal`. Include context in the body when behavior changes. Reference issues as `Fixes #123` when applicable. Pull requests need: a concise summary, validation notes (commands run, screenshots for UI states), and callouts for risk areas. Request review before merging and wait for CI to pass.
