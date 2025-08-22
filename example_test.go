package funcfind_test

import (
	"fmt"

	"github.com/MrUsefull/FuncFind/pkg/funcfind"
)

// ExampleReturning demonstrates basic usage of the Returning function
// to find functions in the fmt package that return an error.
func ExampleReturning() {
	// Find all functions in the fmt package that return an error
	functions, err := funcfind.Returning("fmt", "error")
	if err != nil {
		panic(err)
	}

	numFound := 0
	for fn := range functions {
		fmt.Println(fn.Name())
		numFound++
	}
	fmt.Printf("Found %d function(s) that return error\n", numFound)

	//Output: Errorf
	// Found 1 function(s) that return error
}
