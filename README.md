<!-- This file is safe to edit. Once it exists it will not be overwritten. -->

# exithandler <!-- omit in toc -->

<p align="center">
  <img alt="GitHub Actions" src="https://img.shields.io/github/actions/workflow/status/kilianpaquier/exithandler/integration.yml?branch=main&style=for-the-badge">
  <img alt="GitHub Release" src="https://img.shields.io/github/v/release/kilianpaquier/exithandler?include_prereleases&sort=semver&style=for-the-badge">
  <img alt="GitHub Issues" src="https://img.shields.io/github/issues-raw/kilianpaquier/exithandler?style=for-the-badge">
  <img alt="GitHub License" src="https://img.shields.io/github/license/kilianpaquier/exithandler?style=for-the-badge">
  <img alt="Coverage" src="https://img.shields.io/codecov/c/github/kilianpaquier/exithandler/main?style=for-the-badge">
  <img alt="Go Version" src="https://img.shields.io/github/go-mod/go-version/kilianpaquier/exithandler/main?style=for-the-badge&label=Go+Version">
  <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/kilianpaquier/exithandler?style=for-the-badge">
</p>

---

- [How to use ?](#how-to-use-)
- [Features](#features)

## How to use ?

```sh
go get -u github.com/kilianpaquier/exithandler@latest
```

## Features

The exithandler package exposes two useful functions to handle program terminations:

The first one is `Handle` which will blocked on SIGINT and SIGTERM signals until one of those are sent and then executes the provided function.

```go
func main() {
    // some things to be defined

    go exithandler.Handle(ctx, func(context.Context) {
        // some things to close or execute when the program terminates
    })
}
```

The second one is `HandleFunc` which does the exact same thing, the only difference is that it returns the function which will wait and does not wait directly (as provided in below example).

```go
func main() {
    // some things to be defined

    exithandler := exithandler.HandleFunc(ctx, func(context.Context) {
        // some things to close or execute when the program terminates
    })

    // other things to do

    exithandler()
}
```
