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
	"errors"
	"fmt"
	"log"
	"sort"
	"strings"
	"unicode/utf8"

	"github.com/merenbach/gold-bug/internal/prng"
)

// // Backpermute transforms a string based on a generator function.
// // Backpermute will panic if the transform function returns any invalid string index values.
// func Backpermute(s string, g func() uint) (string, error) {
// 	var out strings.Builder
// 	asRunes := []rune(s)
// 	for range asRunes {
// 		newRune := asRunes[g()]
// 		_, err := out.WriteRune(newRune)
// 		if err != nil {
// 			return "", err
// 		}
// 	}
// 	return out.String(), nil
// }

// Backpermute a rune slice based on a generator function.
// Backpermute will return [E E O H L O] for inputs [H E L L O] and [1 1 4 0 2 4]
// Backpermute will panic if the transform function returns any invalid string index values.
func backpermute(rr []rune, ii []int) ([]rune, error) {
	out := make([]rune, len(ii))

	for n, i := range ii {
		if i < 0 || i >= len(rr) {
			return nil, fmt.Errorf("Index %d out of interval [0, %d)", i, len(rr))
		}
		out[n] = rr[i]
	}

	// for n, i := range out {
	// 	out[n] = rr[ii[n]]
	// }

	return out, nil
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
func Reverse(s string) string {
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
	// Iterate fewer times than the length of padded because we're stepping
	for i := 0; i < len(padded); i += size {
		out = append(out, string(padded[i:i+size]))
	}
	return out
}

// WrapString wraps a string a specified number of indices.
// WrapString will error out if the provided offset is negative.
func WrapString(s string, i int) string {
	// if we simply `return s[i:] + s[:i]`, we're operating on bytes, not runes
	rr := []rune(s)
	return string(rr[i:]) + string(rr[:i])
}

// Affine transform on a string
func Affine(s string, multiply int, add int) (string, error) {
	m := utf8.RuneCountInString(s)

	// TODO: consider using Hull-Dobell satisfaction to determine if `a` is valid (must be coprime with `m`)
	for multiply < 0 {
		multiply += m
	}
	for add < 0 {
		add += m
	}

	lcg := &prng.LCG{
		Modulus:    m,
		Multiplier: 1,
		Increment:  multiply,
		Seed:       add,
	}
	iter, err := lcg.Iterator()
	if err != nil {
		log.Println("Couldn't create LCG")
		return "", err
	}

	positions := make([]int, m)
	for i := range positions {
		positions[i] = iter()
	}

	out, err := backpermute([]rune(s), positions)
	if err != nil {
		return "", err
	}

	return string(out), nil
}

// MakeTranslationTable creates a translation table with source and destination runes, plus optional set of runes to delete.
// MakeTranslationTable is meant in spirit to simulate the Python `str.maketrans` function.
func makeTranslationTable(src []rune, dst []rune, del []rune) (map[rune]rune, error) {
	if len(src) != len(dst) {
		return nil, errors.New("The first two arguments must have equal length")
	}

	t := make(map[rune]rune, len(src))

	// Translate from A to B
	// Indexing here should be safe because of length check above
	for i, r := range src {
		t[r] = dst[i]
	}

	// Mark C for deletion
	for _, r := range del {
		t[r] = (-1)
	}

	return t, nil
}

// MakeTranslationTable creates a translation table with source and destination runes, plus optional set of runes to delete.
// MakeTranslationTable requires all args. The `del` arg may be an empty string.
// MakeTranslationTable is meant in spirit to simulate the Python `str.maketrans` function.
func MakeTranslationTable(src string, dst string, del string) (map[rune]rune, error) {
	return makeTranslationTable([]rune(src), []rune(dst), []rune(del))
}

// Translate a string based on a map of runes.
// Translate returns non-transcodable runes as-is without strict mode.
// Translate will remove any runes that explicitly map to (-1).
func Translate(s string, m map[rune]rune, strict bool) string {
	return strings.Map(func(r rune) rune {
		if o, ok := m[r]; ok {
			// Rune found
			return o
		} else if !strict {
			// Rune not found and strict mode off
			return r
		}
		// Rune not found and strict mode on
		return (-1)
	}, s)
}
