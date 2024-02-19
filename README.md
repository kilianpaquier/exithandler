# exithandler <!-- omit in toc -->

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
