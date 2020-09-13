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
	"strings"
	"unicode"
)

// TODO: Maybe have config object (a la AWS params) to create a table from? Using existing table?

// A Table to hold translation data.
type Table map[rune]rune

// New table.
func New(src string, dst string, del string) (Table, error) {
	m, err := Map(src, dst, del)
	if err != nil {
		return nil, err
	}
	return Table(m), nil
}

// Get a transcoded rune (optionally ignoring case) and a boolean indicating success.
// Get (-1) instead if strict mode is enabled.
// Get the original rune back instead if strict mode is disabled.
func (tt Table) Get(r rune, strict bool, caseless bool) (rune, bool) {
	if o, ok := tt[r]; ok {
		return o, true
	}

	if caseless {
		if o, ok := tt[unicode.ToUpper(r)]; ok {
			return unicode.ToLower(o), true
		} else if o, ok := tt[unicode.ToLower(r)]; ok {
			return unicode.ToUpper(o), true
		}
	}

	if !strict {
		return r, false
	}

	return (-1), false
}

// Map runes in a string.
func (tt Table) Map(s string, strict bool, caseless bool) string {
	return strings.Map(func(r rune) rune {
		o, _ := tt.Get(r, strict, caseless)
		return o
	}, s)
}
