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

import "unicode"

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

// Get a transcoded rune or return (-1) if not found.
func (tt Table) Get(r rune, strict bool, caseless bool) (rune, bool) {
	if o, ok := tt[r]; ok {
		return o, true
	} else if caseless {
		if o, ok := tt[unicode.ToUpper(r)]; ok {
			return unicode.ToLower(o), true
		} else if o, ok := tt[unicode.ToLower(r)]; ok {
			return unicode.ToUpper(o), true
		}
	} else if !strict {
		return r, false
	}
	return (-1), false
}

// // Contains determines if this table contains a rune.
// func (tt Table) Contains(r rune) bool {
// 	if _, ok := tt[r]; ok {
// 		return true
// 	}
// 	return false
// }
