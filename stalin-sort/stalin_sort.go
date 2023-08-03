package stalinsort

import "golang.org/x/exp/constraints"

func Sort[T constraints.Integer](values []T) []T {
	var sorted []T
	for _, value := range values {
		if len(sorted) == 0 || value >= sorted[len(sorted)-1] {
			sorted = append(sorted, value)
		}
	}

	return sorted
}
