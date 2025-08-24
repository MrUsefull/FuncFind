package funcfind

import (
	"go/types"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReturning(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		pkgPath     string
		returnTypes []string
		verify      func(tb testing.TB, fns []*types.Func)
		wantErr     error
	}{
		{
			name:        "Spot check fmt functions for functions that return err",
			pkgPath:     "fmt",
			returnTypes: []string{"error"},
			verify: func(tb testing.TB, found []*types.Func) {
				tb.Helper()
				fnNames := fnNamesFromTypes(found)
				assert.Contains(t, fnNames, "Errorf")
				assert.NotContains(t, fnNames, "Printf")
				assert.NotContains(t, fnNames, "Sprint")
			},
		},
		{
			name:        "Find some functions that return multiple values",
			pkgPath:     "fmt",
			returnTypes: []string{"int", "error"},
			verify: func(tb testing.TB, found []*types.Func) {
				tb.Helper()
				fnNames := fnNamesFromTypes(found)
				assert.Contains(t, fnNames, "Fprintf")
				assert.Contains(t, fnNames, "Fprintln")
				assert.Contains(t, fnNames, "Print")
				assert.NotContains(t, fnNames, "Errorf")
			},
		},
		{
			name:        "Finds no-return functions",
			pkgPath:     "os",
			returnTypes: []string{},
			verify: func(tb testing.TB, found []*types.Func) {
				tb.Helper()
				fnNames := fnNamesFromTypes(found)
				assert.Contains(t, fnNames, "Exit")
			},
		},
		{
			name:    "No matching functions",
			pkgPath: "fmt",
			// order of error, string is chosen because
			// the standard lib always returns error last
			returnTypes: []string{"error", "string"},
			verify: func(tb testing.TB, found []*types.Func) {
				tb.Helper()
				fnNames := fnNamesFromTypes(found)
				assert.Empty(tb, fnNames)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotIter, err := Returning(tt.pkgPath, tt.returnTypes...)
			got := slices.Collect(gotIter)
			require.ErrorIs(t, err, tt.wantErr)
			tt.verify(t, got)
			assertAllExportedFns(t, got)
		})
	}
}

func assertAllExportedFns(tb testing.TB, fns []*types.Func) {
	tb.Helper()

	for _, fn := range fns {
		firstChar := fn.Name()[0:1]
		assert.Equal(tb, strings.ToUpper(firstChar), firstChar, "No unexported functions should be present.")
	}
}

func fnNamesFromTypes(fns []*types.Func) []string {
	fnNames := make([]string, 0, len(fns))
	for _, f := range fns {
		fnNames = append(fnNames, f.Name())
	}
	return fnNames
}
