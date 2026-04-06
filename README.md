# Wire: Automated Initialization in Go

[![Build Status](https://github.com/lorexzer0/wire/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/lorexzer0/wire/actions)
[![godoc](https://godoc.org/github.com/lorexzer0/wire?status.svg)][godoc]

> [!NOTE]
> This is an actively maintained fork of [google/wire](https://github.com/google/wire), which is no longer maintained.
> Maintained by [@lorexzer0](https://github.com/lorexzer0).

Wire is a code generation tool that automates connecting components using
[dependency injection][]. Dependencies between components are represented in
Wire as function parameters, encouraging explicit initialization instead of
global variables. Because Wire operates without runtime state or reflection,
code written to be used with Wire is useful even for hand-written
initialization.

For an overview, see the [introductory blog post][].

[dependency injection]: https://en.wikipedia.org/wiki/Dependency_injection
[introductory blog post]: https://blog.golang.org/wire
[godoc]: https://godoc.org/github.com/lorexzer0/wire
[travis]: https://travis-ci.com/google/wire

## Installing

Install Wire by running:

```shell
go install github.com/lorexzer0/wire/cmd/wire@latest
```

and ensuring that `$GOPATH/bin` is added to your `$PATH`.

## Documentation

- [Tutorial][]
- [User Guide][]
- [Best Practices][]
- [FAQ][]

[Tutorial]: ./_tutorial/README.md
[Best Practices]: ./docs/best-practices.md
[FAQ]: ./docs/faq.md
[User Guide]: ./docs/guide.md

## Project status

This fork picks up where [google/wire](https://github.com/google/wire) left off.
Bug reports, fixes, and feature proposals are welcome — please open an issue or pull request.

## Community

For questions, please use [GitHub Discussions](https://github.com/lorexzer0/wire/discussions).

This project is covered by the Go [Code of Conduct][].

[Code of Conduct]: ./CODE_OF_CONDUCT.md
