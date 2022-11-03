package sliceutil

import (
	"errors"
	"fmt"
	"sort"

	"github.com/merenbach/goldbug/internal/iterutil"
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

// Cycle through a slice. Returns nil if the slice is empty.
func Cycle[T any](xs []T) func() T {
	if len(xs) == 0 {
		return nil
	}

	var cursor int
	return func() T {
		cursor %= len(xs)
		out := xs[cursor]
		cursor++
		return out
	}
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

	g, err := lcg.Iterator()
	if err != nil {
		return nil, fmt.Errorf("couldn't initialize LCG: %w", err)
	}

	ys, err := Backpermute(xs, iterutil.Take(m, g))
	if err != nil {
		return nil, fmt.Errorf("couldn't backpermute input: %w", err)
	}

	return ys, nil
}

// Filter slice values through a function.
func Filter[T any](xs []T, f func(T) bool) []T {
	out := make([]T, 0)
	for _, o := range xs {
		if f(o) {
			out = append(out, o)
		}
	}
	return out
}

// Keyword transform a slice.
func Keyword[T comparable](xs []T, keyword []T) []T {
	set := make(map[T]struct{})
	for _, x := range xs {
		if _, ok := set[x]; !ok {
			set[x] = struct{}{}
		}
	}

	// Retain in keyword only elements shared with xs
	filteredKeyword := Filter(keyword, func(o T) bool {
		_, ok := set[o]
		return ok
	})

	filteredKeyword = append(filteredKeyword, xs...)
	return Deduplicate(filteredKeyword)
}

// Map slice values through a function.
func Map[T any, U any](xs []T, f func(T) U) []U {
	out := make([]U, len(xs))
	for i, o := range xs {
		out[i] = f(o)
	}
	return out
}

// Zipmap creates a new map by zipping a key scalar with a slice of values.
// Zipmap requires both parameters to have the same length.
func Zipmap[T comparable, U any](xs []T, ys []U) (map[T]U, error) {
	if len(xs) != len(ys) {
		return nil, errors.New("both parameters must have the same length")
	}

	m := make(map[T]U)
	for i, x := range xs {
		m[x] = ys[i]
	}

	return m, nil
}

/// Zigzag sequence, of primary use in the rail fence cipher.
/// The period is the length of the sequence before any repetition would occur.
/// A single period will be returned.
// pub fn zigzag(period: usize) -> Vec<usize> {
//     (0..period).map(|n| cmp::min(n, period - n)).collect()
// }
