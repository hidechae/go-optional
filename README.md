[![Test](https://github.com/hidechae/go-optional/actions/workflows/test.yml/badge.svg)](https://github.com/hidechae/go-optional/actions/workflows/test.yml)

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

// Get value with error handling
value, err := opt.Get()
if err != nil {
    // Handle error (returns optional.ErrGetValueFromNone if None)
}

// Get value with fallback
value := opt.GetOr(0) // Returns 42, or 0 if None

// Convert to pointer (returns nil if None)
ptr := opt.ToPtr()
```

### Functional Operations

```go
opt := optional.Some(42)

// Map: Transform the value if present
doubled := optional.Map(opt, func(x int) int {
    return x * 2
}) // Some(84)

// Map with type conversion
str := optional.Map(opt, func(x int) string {
    return fmt.Sprintf("Value: %d", x)
}) // Some("Value: 42")

// Map on None returns None
noneOpt := optional.None[int]()
result := optional.Map(noneOpt, func(x int) int {
    return x * 2
}) // None
```

### JSON Marshalling/Unmarshalling

The Option type supports JSON marshalling and unmarshalling. `Some` values are marshalled as the contained value, while `None` values are marshalled as `null`.

```go
type User struct {
    Name  string         `json:"name"`
    Email Option[string] `json:"email"`
    Age   Option[int]    `json:"age"`
}

user := User{
    Name:  "Alice",
    Email: optional.Some("alice@example.com"),
    Age:   optional.None[int](),
}

// Marshal to JSON
data, err := json.Marshal(user)
// Result: {"name":"Alice","email":"alice@example.com","age":null}

// Unmarshal from JSON
jsonData := `{"name":"Bob","email":"bob@example.com","age":30}`
var newUser User
err = json.Unmarshal([]byte(jsonData), &newUser)
// newUser.Email will be Some("bob@example.com")
// newUser.Age will be Some(30)

// Null values become None
jsonData2 := `{"name":"Charlie","email":null,"age":null}`
var user2 User
err = json.Unmarshal([]byte(jsonData2), &user2)
// user2.Email will be None[string]()
// user2.Age will be None[int]()
```

The JSON marshalling also works with the `omitzero` tag to omit `None` values from JSON output:

```go
type Config struct {
    Debug Option[bool] `json:"debug,omitzero"`
}

config := Config{Debug: optional.None[bool]()}
data, _ := json.Marshal(config)
// Result: {} (empty object, debug field omitted)
```

## API Reference

### Types

- `Option[T any]` - Generic optional type that can hold any type T

### Variables

- `ErrGetValueFromNone` - Error returned when trying to get value from None

### Functions

- `Some[T any](v T) Option[T]` - Create an Option containing the given value
- `None[T any]() Option[T]` - Create an empty Option
- `FromPtr[T any](v *T) Option[T]` - Create an Option from a pointer (nil becomes None)
- `Map[T, V any](o Option[T], f func(T) V) Option[V]` - Transform the value if present

### Methods

- `IsSone() bool` - Returns true if the Option contains a value
- `IsNone() bool` - Returns true if the Option is empty
- `Get() (T, error)` - Returns the contained value or an error if None
- `GetOr(fallback T) T` - Returns the contained value or the fallback if None
- `ToPtr() *T` - Returns a pointer to the contained value or nil if None

## Running Tests

```bash
go test
```

## License

This project is open source and available under the MIT License.