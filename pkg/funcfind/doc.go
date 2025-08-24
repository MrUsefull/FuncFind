/*
Package funcfind provides static analysis tools for discovering Go functions by their return types.

# Overview

FuncFind uses Go's type system and the golang.org/x/tools/go/packages library to analyze
Go packages and find functions that return specific types. This is particularly useful for
static analysis tools, code generators, and developer utilities that need to understand
the structure of Go codebases.

# Key Features

  - Type-based function discovery using string matching
  - Support for functions with variable-length return types
  - Lazy evaluation through Go 1.23+ iterators
  - Support for any accessible Go package
  - Memory-efficient processing of large codebases
  - Early termination support for performance

# Usage Patterns

The primary function is Returning, which takes a package path and one or more return type strings:

	// Find functions returning a single error
	functions, err := funcfind.Returning("fmt", "error")
	if err != nil {
		return err
	}

	for fn := range functions {
		fmt.Printf("Function %s returns error\n", fn.Name())
	}

	// Find functions returning (int, error)
	functions, err = funcfind.Returning("fmt", "int", "error")
	if err != nil {
		return err
	}

	for fn := range functions {
		fmt.Printf("Function %s returns (int, error)\n", fn.Name())
	}

# Type Matching

Function signatures are matched by comparing the string representation of their return types.
Functions must match the exact number and order of return types specified. For example:

  - func FnName() error            // ✓ matches Returning(pkg, "error")
  - func FnName1() (int, error)    // ✓ matches Returning(pkg, "int", "error")
  - func FnName2() (error, int)    // ✗ wrong order for Returning(pkg, "int", "error")
  - func FnName5()                 // ✓ matches Returning(pkg)
  - func FnName4() string          // ✗ different type for Returning(pkg, "error")

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
  - Type matching is string comparison, order matters
  - No support for interface satisfaction matching

# Error Handling

Errors are returned when:
  - Package cannot be loaded (e.g., doesn't exist, compilation errors)
  - Package path is malformed
  - Module resolution fails
*/
package funcfind
