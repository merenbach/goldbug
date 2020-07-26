// Copyright 2019 Andrew Merenbach
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

package railfence

import (
	"strings"
	"unicode/utf8"
)

// A Cipher implements the rail fence (or zig-zag) cipher.
type Cipher struct {
	Rails int
}

// Encipher a message.
func (c *Cipher) Encipher(s string) string {
	if c.Rails == 1 {
		return s
	}

	cc := make(grid, utf8.RuneCountInString(s))
	cc.renumber(c.Rails)

	cc.fill(s)
	cc.sortByRow()

	return cc.String()
}

// Decipher a message.
func (c *Cipher) Decipher(s string) string {
	if c.Rails == 1 {
		return s
	}

	cc := make(grid, utf8.RuneCountInString(s))
	cc.renumber(c.Rails)

	cc.sortByRow()
	cc.fill(s)
	cc.sortByCol()

	return cc.String()
}

// // Cycle length for cipher.
// func (c *Cipher) Cycle() int {
// 	return 2 * (c.Rails - 1)
// }

// Tableau output for this cipher.
func (c *Cipher) Tableau(s string) string {
	if c.Rails == 1 {
		return s
	}

	cc := make(grid, utf8.RuneCountInString(s))
	cc.renumber(c.Rails)

	cc.fill(s)
	cc.sortByRow()

	var currentRow int
	var currentCol int

	var out strings.Builder
	for _, c := range cc {
		if c.Row > currentRow {
			currentRow = c.Row
			currentCol = 0
			out.WriteRune('\n')
		}

		for c.Col > currentCol {
			out.WriteRune(' ')
			currentCol++
		}
		currentCol++

		out.WriteRune(c.Rune)
	}
	return out.String()
}
