# Changelog

## v0.8.3

### New Features

- **`-wrap_errors` flag**: Opt-in flag for `wire gen` and `wire diff` that wraps
  provider errors with the provider name. When enabled, generated error returns
  change from `return ..., err` to `return ..., fmt.Errorf("providerName: %w", err)`.
  This makes it easier to identify which provider failed in a dependency chain.

  Usage:
  ```shell
  wire gen -wrap_errors ./...
  ```

  **Breaking change when enabled**: Code that compares provider errors using
  direct equality (`err == ErrFoo`) or exact string matching
  (`err.Error() == "message"`) will break. Use `errors.Is` and `errors.As`
  instead, which work correctly with wrapped errors.

### Improvements

- Generated output now includes source line numbers as comments above each
  copied non-injector declaration (e.g. `// foo.go:28`).
- Error messages for inaccessible value expressions now report the position of
  the value expression itself, not the injector function.
- Replaced hardcoded command map with dynamic discovery via `VisitCommands`.
- `wire.Bind` now validates that arguments are type expressions (`new(T)` or
  `(*T)(nil)`). Passing variables, function calls, or other computed values
  will now produce a clear error instead of potentially confusing type errors.
