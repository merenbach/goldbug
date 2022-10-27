// Copyright 2018 Andrew Merenbach
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stringutil

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"

	"golang.org/x/exp/constraints"
)

// Backpermute a data slice based on a slice of index values.
// Backpermute will return [E E O H L O] for inputs [H E L L O] and [1 1 4 0 2 4]
// Backpermute will return an error if the transform function returns any invalid string index values.
func Backpermute[T any](xs []T, by []int) ([]T, error) {
	var ys []T
	for _, i := range by {
		if i < 0 || i >= len(xs) {
			return nil, fmt.Errorf("slice index %d out of range [0, %d)", i, len(xs))
		}
		ys = append(ys, xs[i])
	}
	return ys, nil
}

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

// Key a string with prefix text.
// TODO: rename Prefix? or something to allow semantically for suffix counterpart?
func Key(s string, k string) string {
	return Deduplicate(k + s)
}

// Deduplicate removes recurrences for runes from a string, preserving order of first appearance.
func Deduplicate(s string) string {
	seen := make(map[rune]struct{})
	return strings.Map(func(r rune) rune {
		if _, ok := seen[r]; !ok {
			seen[r] = struct{}{}
			return r
		}
		return (-1)
	}, s)
}

// Intersect removes runes from a string if they don't occur in another string.
func intersect(s, charset string) string {
	seen := make(map[rune]bool)

	for _, r := range charset {
		seen[r] = true
	}

	return strings.Map(func(r rune) rune {
		if _, ok := seen[r]; ok {
			return r
		}
		return -1
	}, s)
}

// Reverse the order of runes in a string.
/*func Reverse(s string) string {
	r := []rune(s)
	sort.SliceStable(r, func(i, j int) bool {
		return true
	})
	return string(r)
}*/

// Chunk divides a string into groups.
func chunk(s string, size int, delimiter rune) string {
	return strings.Join(groupString(s, size, 'X'), string(delimiter))
}

// DiffToMod returns the difference between a and the nearest multiple of m.
func diffToMod(a, m int) int {
	if remainder := a % m; remainder != 0 {
		return m - remainder
	}
	return 0
}

// GroupString divides a string into groups.
func groupString(s string, size int, padding rune) []string {
	out := make([]string, 0)
	nullCount := diffToMod(utf8.RuneCountInString(s), size)
	nulls := strings.Repeat(string(padding), nullCount)
	padded := []rune(s + nulls)
	// Iterate fewer times than the length of padded because we're stepping
	for i := 0; i < len(padded); i += size {
		out = append(out, string(padded[i:i+size]))
	}
	return out
}

// // WrapString wraps a string a specified number of indices.
// // WrapString will error out if the provided offset is negative.
// func WrapString(s string, i int) string {
// 	// if we simply `return s[i:] + s[:i]`, we're operating on bytes, not runes
// 	// When adapting for slices in the input, remember to copy so as to not modify the original
// 	rr := []rune(s)
// 	return string(append(rr[i:], rr[:i]...))
// 	// return string(rr[i:]) + string(rr[:i])
// }
