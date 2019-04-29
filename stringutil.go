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

package main

import (
	"sort"
	"strings"
	"unicode/utf8"
)

// Backpermute transforms a string based on a generator function.
// Backpermute will panic if the transform function returns any invalid string index values.
func backpermute(s string, g func() uint) (string, error) {
	var out strings.Builder
	asRunes := []rune(s)
	for range asRunes {
		newRune := asRunes[g()]
		_, err := out.WriteRune(newRune)
		if err != nil {
			return "", err
		}

	}
	return out.String(), nil
}

// Deduplicate removes recurrences for runes from a string, preserving order of first appearance.
func deduplicate(s string) string {
	seen := make(map[rune]bool)
	return strings.Map(func(r rune) rune {
		if _, ok := seen[r]; !ok {
			seen[r] = true
			return r
		}
		return -1
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

// ReverseString reverses the runes in a string.
func reverseString(s string) string {
	r := []rune(s)
	sort.SliceStable(r, func(i, j int) bool {
		return true
	})
	return string(r)
}

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
	for i := 0; i < len(padded); i += size {
		out = append(out, string(padded[i:i+size]))
	}
	return out
}

// WrapString wraps a string a specified number of indices.
// WrapString will error out if the provided offset is negative.
func wrapString(s string, i int) string {
	// if we simply `return s[i:] + s[:i]`, we're operating on bytes, not runes
	rr := []rune(s)
	return string(rr[i:]) + string(rr[:i])
}
