package sort

import "golang.org/x/exp/constraints"

type Sort[T constraints.Integer] interface {
	Sort(values []T) []T
}

type SortFn[T constraints.Integer] func(values []T) []T
