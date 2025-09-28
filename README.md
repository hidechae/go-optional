# go-optional

A Go library providing an optional type for handling nullable values in a type-safe way.

## Overview

This library implements the Option pattern (similar to Rust's `Option` or Haskell's `Maybe`) for Go using generics. It helps eliminate null pointer exceptions and makes the presence or absence of values explicit in your type system.

## Installation

```bash
go get github.com/hidechae/go-optional
```

## Usage

### Creating Options

```go
package main

import "github.com/hidechae/go-optional"

// Create an Option with a value
opt1 := optional.Some(42)

// Create an empty Option
opt2 := optional.None[int]()

// Create from a pointer (nil pointer becomes None)
var ptr *int = nil
opt3 := optional.FromPtr(ptr)
```

### Checking Option State

```go
opt := optional.Some(42)

if opt.IsSone() {
    // Option contains a value
}

if opt.IsNone() {
    // Option is empty
}
```

### Getting Values

```go
opt := optional.Some(42)

// Get value with fallback
value := opt.GetOr(0) // Returns 42, or 0 if None

// Convert to pointer (returns nil if None)
ptr := opt.ToPtr()
```

## API Reference

### Types

- `Option[T any]` - Generic optional type that can hold any type T

### Functions

- `Some[T any](v T) Option[T]` - Create an Option containing the given value
- `None[T any]() Option[T]` - Create an empty Option
- `FromPtr[T any](v *T) Option[T]` - Create an Option from a pointer (nil becomes None)

### Methods

- `IsSone() bool` - Returns true if the Option contains a value
- `IsNone() bool` - Returns true if the Option is empty
- `GetOr(fallback T) T` - Returns the contained value or the fallback if None
- `ToPtr() *T` - Returns a pointer to the contained value or nil if None

## Running Tests

```bash
go test
```

## License

This project is open source and available under the MIT License.