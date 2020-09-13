// Copyright 2020 Andrew Merenbach
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

package translation

import (
	"errors"
	"strings"
)

// MakeMap maps source runes to destination runes.
// MakeMap maps to (-1) any runes to delete.
// MakeMap is modeled after the Python `str.maketrans` function.
func makeMap(a string, b string, c string) (map[rune]rune, error) {
	src, dst, del := []rune(a), []rune(b), []rune(c)

	if len(src) != len(dst) {
		return nil, errors.New("The first two arguments must have equal length")
	}

	t := make(map[rune]rune, len(src))

	// Translate from A to B
	for i, r := range src {
		t[r] = dst[i]
	}

	// Mark C for deletion
	for _, r := range del {
		t[r] = (-1)
	}

	return t, nil
}

// Translate a string based on a map of runes.
// Translate returns non-transcodable runes as-is without strict mode.
// Translate will remove any runes that explicitly map to (-1).
func translate(s string, m map[rune]rune, strict bool) string {
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
