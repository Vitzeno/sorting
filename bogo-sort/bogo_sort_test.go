package bogosort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBogoSort(t *testing.T) {
	for _, tc := range []struct {
		unordered []int
		ordered   []int
	}{
		{
			unordered: []int{3, 1, 2},
			ordered:   []int{1, 2, 3},
		},
	} {
		assert.Equal(t, tc.ordered, Sort(tc.unordered))
	}
}
