package sleepsort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSleepSort(t *testing.T) {
	for _, tc := range []struct {
		unordered []int
		ordered   []int
	}{
		{
			unordered: []int{3, 2, 4, 5, 1},
			ordered:   []int{1, 2, 3, 4, 5},
		},
	} {
		assert.Equal(t, Sort(tc.unordered), tc.ordered)
	}
}
