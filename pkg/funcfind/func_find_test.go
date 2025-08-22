package funcfind

import (
	"go/types"
	"iter"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReturning(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		pkgPath    string
		returnType string
		verify     func(tb testing.TB, fns iter.Seq[*types.Func])
		wantErr    error
	}{
		{
			name:       "Spot check fmt functions for functions that return err",
			pkgPath:    "fmt",
			returnType: "error",
			verify: func(tb testing.TB, s iter.Seq[*types.Func]) {
				tb.Helper()
				found := slices.Collect(s)
				fnNames := make([]string, 0, len(found))
				for _, f := range found {
					fnNames = append(fnNames, f.Name())
				}
				assert.Contains(t, fnNames, "Errorf")
				assert.NotContains(t, fnNames, "Printf")
				assert.NotContains(t, fnNames, "Sprint")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Returning(tt.pkgPath, tt.returnType)
			require.ErrorIs(t, err, tt.wantErr)
			tt.verify(t, got)
		})
	}
}
