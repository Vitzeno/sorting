package stalinsort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStalinSort(t *testing.T) {
	for _, tc := range []struct {
		unordered []int
		ordered   []int
	}{
		{
			unordered: []int{3, 2, 4, 5, 1},
			ordered:   []int{3, 4, 5},
		},
		{
			unordered: []int{4, 6, 7, 2, 3, 1, 9},
			ordered:   []int{4, 6, 7, 9},
		},
		{
			unordered: []int{9, 5, 2, 6, 7, 1},
			ordered:   []int{9},
		},
	} {
		assert.Equal(t, tc.ordered, Sort(tc.unordered))
	}
}
