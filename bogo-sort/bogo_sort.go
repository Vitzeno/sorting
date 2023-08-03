package bogosort

import (
	"math/rand"
	"sort"

	"golang.org/x/exp/constraints"
)

func Sort[T constraints.Integer](values []T) []T {
	isSorted := false
	for !isSorted {
		rand.Shuffle(len(values), func(i, j int) { values[i], values[j] = values[j], values[i] })
		isSorted = sort.SliceIsSorted(values, func(i, j int) bool { return values[i] < values[j] })
	}

	return values
}
