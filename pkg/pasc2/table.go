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

package pasc2

import (
	"unicode"

	"github.com/merenbach/goldbug/internal/translation"
)

// A Table to hold translation data.
type Table map[rune]translation.Table

// NewTable creates a new table.
// func NewTable(src string, dst string, del string) (Table, error) {
// 	m, err := makeMap(src, dst, del)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return m, nil
// }

// Get a transcoded rune (optionally ignoring case) and a boolean indicating success.
// Get (-1) instead if strict mode is enabled.
// Get the original rune back instead if strict mode is disabled.
func (tt Table) Get(r rune, k rune, strict bool, caseless bool) (rune, bool) {
	m, ok := tt[k]

	if caseless {
		m, ok = tt[unicode.ToUpper(r)]
		if !ok {
			m, ok = tt[unicode.ToLower(r)]
		}
	}

	return m.Get(r, strict, caseless)

	// if !strict {
	// 	return r, false
	// }

	// return (-1), false
}
