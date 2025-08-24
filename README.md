# FuncFind

[![CI](https://github.com/MrUsefull/FuncFind/workflows/CI/badge.svg)](https://github.com/MrUsefull/FuncFind/actions/workflows/ci.yml)

FuncFind is a Go library for discovering functions by their return types using static analysis. It leverages Go's type system and the `go/types` package to efficiently find functions that return specific types across Go packages.

## Features

- **Type-based function discovery**: Find all functions that return specific type(s)
- **Lazy evaluation**: Uses Go 1.23+ iterators for memory-efficient traversal
- **Flexible matching**: Finds functions matching an exact sequence of return types

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
    
    // Find all functions in the fmt package that return (int, error)
    functions, err = funcfind.Returning("fmt", "int", "error")
    if err != nil {
        panic(err)
    }
    
    for fn := range functions {
        fmt.Printf("%s returns (int, error)\n", fn.Name())
    }
}
```

## Limitations

- Functions must match the exact number and order of return types specified
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
