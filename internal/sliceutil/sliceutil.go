package sliceutil

import (
	"errors"
	"fmt"
	"sort"

	"github.com/merenbach/goldbug/internal/mathutil"
	"github.com/merenbach/goldbug/internal/prng"
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

// Affine transform a slice and return the result.
func Affine[T any](xs []T, slope int, intercept int) ([]T, error) {
	m := len(xs)

	if m == 0 {
		return xs, nil
	}

	for slope < 0 {
		slope += m
	}
	for intercept < 0 {
		intercept += m
	}

	if !mathutil.Coprime(m, slope) {
		return nil, errors.New("slope and string length must be coprime")
	}

	lcg := &prng.LCG{
		Modulus:    m,
		Multiplier: 1,
		Increment:  slope,
		Seed:       intercept,
	}

	positions, err := lcg.Slice(m)
	if err != nil {
		return nil, fmt.Errorf("couldn't initialize LCG: %w", err)
	}

	ys, err := Backpermute(xs, positions)
	if err != nil {
		return nil, fmt.Errorf("couldn't backpermute input: %w", err)
	}

	return ys, nil
}

// Keyword transform a slice.
func Keyword[T comparable](xs []T, keyword []T) ([]T, error) {
	set := make(map[T]struct{})
	for _, x := range xs {
		if _, ok := set[x]; !ok {
			set[x] = struct{}{}
		}
	}

	// Filter out anything from the keyword that isn't in the primary slice
	filteredKeyword := make([]T, 0)
	for _, k := range keyword {
		if _, ok := set[k]; ok {
			filteredKeyword = append(filteredKeyword, k)
		}
	}

	filteredKeyword = append(filteredKeyword, xs...)
	return Deduplicate(filteredKeyword), nil
}
