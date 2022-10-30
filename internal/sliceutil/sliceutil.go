package sliceutil

import (
	"fmt"
	"sort"

	"golang.org/x/exp/constraints"
)

/// Argsort returns the indices that would sort an array.
/// The naming is based on the numpy extension's name for this concept.
///
/// <https://numpy.org/doc/stable/reference/generated/numpy.argsort.html>
///
/// This is effectively a Schwartzian transform or decorate-sort-undecorate.
///
///   1. Attach numbers to each item in the collection.
///   2. Rearrange the collection such that it is now sorted lexically. This will scramble the numbers.
///   3. Return only the numbers now.
func Argsort[T constraints.Ordered](xs []T) []int {
	ys := make([]int, len(xs))
	for i := range ys {
		ys[i] = i
	}

	sort.SliceStable(ys, func(i int, j int) bool {
		return xs[ys[i]] < xs[ys[j]]
	})

	return ys
}

// Backpermute a data slice based on a slice of index values.
// Backpermute will return [E E O H L O] for inputs [H E L L O] and [1 1 4 0 2 4]
// Backpermute will return an error if the transform function returns any invalid string index values.
func Backpermute[T any](xs []T, by []int) ([]T, error) {
	ys := make([]T, 0)
	for _, i := range by {
		if i < 0 || i >= len(xs) {
			return nil, fmt.Errorf("slice index %d out of range [0, %d)", i, len(xs))
		}
		ys = append(ys, xs[i])
	}
	return ys, nil
}

// Deduplicate removes recurrences for elements from a sequence.
// Deduplicate is stable in that it preserves the order of first appearance.
func Deduplicate[T comparable](xs []T) []T {
	seen := make(map[T]struct{})
	out := make([]T, 0)
	for _, x := range xs {
		if _, ok := seen[x]; !ok {
			seen[x] = struct{}{}
			out = append(out, x)
		}
	}
	return out
}
