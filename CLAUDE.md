# Goztl: Go client library for the Zentral API

Goztl wraps the REST API exposed by [Zentral](https://github.com/zentralopensource/zentral), a Django platform that orchestrates MDM, Munki, Osquery and Santa to manage Apple devices. This repo is the Go client only — server-side concerns (event pipeline, agents, Django apps) live in the Zentral repo.

Module path: `github.com/zentralopensource/goztl`. Single package `goztl` at the repo root. `go.mod` pins Go 1.24; CI matrix is Go 1.25.x on ubuntu/macos/windows.

## Local environment: Docker only

We develop on macOS machines that run a binary blocker. **Do not install or invoke a local `go` toolchain.** Every Go command — build, test, format, `mod tidy`, `generate` — must run inside a container.

One-shot pattern (use the same image tag CI uses):

```
docker run --rm -v "$PWD":/src -w /src golang:1.25 go test -race ./...
docker run --rm -v "$PWD":/src -w /src golang:1.25 gofmt -s -l .
docker run --rm -v "$PWD":/src -w /src golang:1.25 go mod tidy
docker run --rm -v "$PWD":/src -w /src golang:1.25 go generate -x ./...
```

The repo also ships a devcontainer (`.devcontainer/devcontainer.json`, `mcr.microsoft.com/devcontainers/go:1.24`) — open the folder in a devcontainer-aware editor to get an interactive shell with the toolchain available inside the container.

If a command would shell out to `go` or `gofmt` directly on the host, rewrite it as a `docker run …` invocation before executing.

## Code layout

The package is flat: one `<area>_<resource>.go` file per Zentral API resource with a sibling `_test.go`. Areas mirror Zentral's modules: `mdm_*`, `monolith_*`, `munki_*`, `osquery_*`, `santa_*`, `probes_*`, `gws_*` (Google Workspace), plus inventory primitives (`tags.go`, `taxonomies.go`, `meta_business_units.go`, `jmespath_checks.go`) and shared infrastructure:

- `goztl.go` — `Client` struct, `NewClient`, request/response plumbing, pagination (`resolveAllPages[T]`), `ListOptions`, `addOptions`, pointer helpers (`Int`, `String`).
- `common.go` — shared payload types (`HTTPHeader`, `EventFilter`, `EventFilterSet`).
- `enrollment_secret.go` — shared `EnrollmentSecret` / `EnrollmentSecretRequest` reused by every agent enrollment resource.
- `errors.go` — `ArgError` / `NewArgError` for input validation.
- `timestamp.go` — `Timestamp` type with ISO8601 + Unix unmarshalling.
- `strings.go` — `Stringify` reflection-based pretty-printer used by every resource's `String()` method.
- `goztl_test.go` — `setup()` httptest harness + `testMethod` / `testHeader` / `testQueryArg` / `testBody` helpers.

## Adding a new resource

Each resource follows the same shape — copy a close neighbour (e.g. [monolith_catalogs.go](monolith_catalogs.go) for an int-keyed resource, [probes_actions.go](probes_actions.go) for a string-keyed one) and adapt:

1. Declare `const xxxBasePath = "area/resource/"` (trailing slash, no leading slash — paths resolve against `Client.BaseURL`).
2. Define the `XxxService` interface — typically `List / GetByID / GetByName / Create / Update / Delete`, each returning `(value, *Response, error)`.
3. Define the `XxxServiceOp` struct holding `client *Client` and assert `var _ XxxService = &XxxServiceOp{}`.
4. Define the response struct `Xxx` (with `json:"…"` tags matching Zentral) and a `XxxRequest` struct for Create/Update bodies. Add `func (x Xxx) String() string { return Stringify(x) }`.
5. Define a `listXxOptions` struct with `url:"name,omitempty"` tags for any server-side filters used by `GetByName`-style lookups.
6. Implement methods. Validation goes through `NewArgError`. List methods funnel through a private `list(ctx, opt, xOpt)` that calls `addOptions` twice (once for `*ListOptions`, once for filter options) then `resolveAllPages[Xxx](ctx, s.client, path)` — never page manually.
7. Wire the service into [goztl.go](goztl.go): add a field on `Client`, a typed-interface declaration, and an initializer line in `NewClient` (keep the alphabetical-within-area grouping).
8. Add a `<resource>_test.go` next to it. Drive each method against an `httptest` mux registered in `setup()`; assert request method/headers/query/body with the helpers in [goztl_test.go](goztl_test.go); compare returned values with `cmp.Equal` and request bodies with `assert.Equal` from `testify`.

## Conventions

Carry these over from the upstream Zentral CLAUDE.md:

- Don't comment self-explanatory code. Well-named identifiers say what the code does — comments are for **why** something exists.
- Match existing style. JSON tag casing, request/response struct splits, pointer vs value for optional fields (`*string` for nullable, plain `string` for required), and the `XxxRequest` naming are consistent across the codebase; new resources should not invent variants.

## CI gates

[.github/workflows/test.yml](.github/workflows/test.yml) runs on every push/PR to `main`:

- `go generate -x ./...` produces a zero diff.
- `gofmt -s -l .` is empty (ubuntu only).
- `go test -v -race ./...` passes on ubuntu / macos / windows with Go 1.25.x.

Run the equivalent Docker commands locally before pushing.

Releases are tag-driven (`v*`) via [GoReleaser](.goreleaser.yaml) — `builds.skip: true`, so this is a library-only release that just attaches a changelog/source archive.
