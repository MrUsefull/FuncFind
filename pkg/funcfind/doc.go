/*
Package funcfind provides static analysis tools for discovering Go functions by their return types.

# Overview

FuncFind uses Go's type system and the golang.org/x/tools/go/packages library to analyze
Go packages and find functions that return specific types. This is particularly useful for
static analysis tools, code generators, and developer utilities that need to understand
the structure of Go codebases.

# Key Features

  - Type-based function discovery using string matching
  - Lazy evaluation through Go 1.23+ iterators
  - Support for any accessible Go package
  - Memory-efficient processing of large codebases
  - Early termination support for performance

# Usage Patterns

The primary function is Returning, which takes a package path and a return type string:

	functions, err := funcfind.Returning("fmt", "error")
	if err != nil {
		return err
	}

	for fn := range functions {
		fmt.Printf("Function %s returns error\n", fn.Name())
	}

# Type Matching

Function signatures are matched by comparing the string representation of their return types.
Only functions with exactly one return value are considered. For example:

  - func() error          ✓ matches "error"
  - func() (error, bool)  ✗ multiple returns
  - func()                ✗ no returns
  - func() string         ✗ different type

# Performance Considerations

The iterator pattern allows for efficient memory usage and early termination:

	// Stop after finding the first match
	for fn := range functions {
		if fn.Name() == "target" {
			break // Iterator stops here
		}
	}

# Limitations

  - Only exported functions are found in external packages
  - Package must be accessible and compilable
  - Type matching is exact string comparison
  - No support for interface satisfaction matching

# Error Handling

Errors are returned when:
  - Package cannot be loaded (e.g., doesn't exist, compilation errors)
  - Package path is malformed
  - Module resolution fails

All errors are wrapped with context to aid debugging.
*/
package funcfind