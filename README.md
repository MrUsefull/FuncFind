# FuncFind

[![CI](https://github.com/MrUsefull/FuncFind/workflows/CI/badge.svg)](https://github.com/MrUsefull/FuncFind/actions/workflows/ci.yml)

FuncFind is a Go library for discovering functions by their return types using static analysis. It leverages Go's type system and the `go/types` package to efficiently find functions that return specific types across Go packages.

## Features

- **Type-based function discovery**: Find all functions that return a specific type
- **Lazy evaluation**: Uses Go 1.23+ iterators for memory-efficient traversal
- **Standard library compatible**: Works with any Go package, including the standard library
- **Exact matching**: Finds functions with exactly one return value of the specified type

## Installation

```bash
go get github.com/MrUsefull/FuncFind
```

## Usage

### Basic Example

```go
package main

import (
    "fmt"
    "slices"
    
    "github.com/MrUsefull/FuncFind/pkg/funcfind"
)

func main() {
    // Find all functions in the fmt package that return an error
    functions, err := funcfind.Returning("fmt", "error")
    if err != nil {
        panic(err)
    }
    
    for fn := range functions {
        fmt.Println(fn.Name())
    }
}
```

## How It Works

FuncFind uses Go's `golang.org/x/tools/go/packages` to load package information and `go/types` for type analysis. The process:

## Limitations

- Only finds functions with exactly **one** return value.
  - It's possible to extend to arbitrary number of return values later.
- Type matching is done by string comparison of type representations
- Requires the target package to be accessible and compilable
- Does not find unexported functions in external packages

## Development

### Running Tests

```bash
go test ./...
```

### Running Linting

```bash
golangci-lint run
```

### Building

```bash
go build ./...
```

## Contributing

Contributions are welcome! Please feel free to submit issues or pull requests.

## License

MIT
