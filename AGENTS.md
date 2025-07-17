# AGENTS.md

## Build, Lint, and Test
- **Build:** `make build` or `go build --tags "fts5" ./cmd/server`
- **Lint:** `go fmt ./...` and `go vet ./...` (no custom linter config)
- **Test:** No *_test.go files present, but standard Go: `go test ./...`
  - To run a single test: `go test ./path/to/package -run TestName`

## Code Style Guidelines
- Use standard Go formatting (`gofmt`) and idiomatic import grouping: stdlib, external, local.
- Types: Prefer explicit types; use struct tags for DB mapping.
- Naming: CamelCase for exported types/functions, lowerCamelCase for locals.
- Error handling: Always check errors; log or return as appropriate.
- Imports: Group and order as per Go conventions.
- Use context-appropriate error messages.
- Use `sqlx` for DB, Echo for HTTP, templates for rendering.
- No custom lint/staticcheck config; follow Go best practices.


## Hypermedia & Client Architecture
- The application follows HATEOAS (Hypermedia as the Engine of Application State) principles for API and UI design.
- Client-side navigation and partial updates use Alpine AJAX (see templates/head.html), enabling dynamic, hypermedia-driven interactions.

No Cursor or Copilot rules detected. Follow idiomatic Go and project conventions.