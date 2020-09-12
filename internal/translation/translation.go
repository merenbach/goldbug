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

import "strings"

type T2 map[rune]rune

// New T2.
func New(src string, dst string, del string) (T2, error) {
	m, err := makeMap(src, dst, del)
	if err != nil {
		return nil, err
	}
	return T2(m), nil
}

// Translate a string.
func (tt T2) Translate(s string, strict bool) string {
	return strings.Map(func(r rune) rune {
		if o, ok := tt[r]; ok {
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

// A Table to hold translation data.
type Table struct {
	Src string
	Dst string
	Del string

	Strict bool
}

// Map source runes to destination runes and map to (-1) any runes to delete.
func (tt *Table) Map() (map[rune]rune, error) {
	return makeMap(tt.Src, tt.Dst, tt.Del)
}

// Translate a string based on a map of runes.
// Translate returns non-transcodable runes as-is without strict mode.
// Translate will remove any runes that explicitly map to (-1).
func (tt *Table) Translate(s string) (string, error) {
	m, err := tt.Map()
	if err != nil {
		return "", err
	}
	return translate(s, m, tt.Strict), nil
}
