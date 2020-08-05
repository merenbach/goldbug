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
	"unicode/utf8"

	"github.com/merenbach/goldbug/internal/grid"
)

// Cipher implements a rail fence (or zig-zag) cipher.
type Cipher struct {
	Rows int
}

// Row for the message character at the given index.
func (c *Cipher) row(i int) int {
	k := c.Rows - 1
	return k - abs(k-i%(2*k))
}

// Makegrid creates a grid and numbers its cells.
func (c *Cipher) makegrid(n int) grid.Grid {
	g := make(grid.Grid, n)

	for i := range g {
		g[i].Col = i

		// Cycle length is 2*(rows - 1), or 2*k
		g[i].Row = c.row(i)
	}

	return g
}

// Encipher a message.
func (c *Cipher) Encipher(s string) (string, error) {
	if c.Rows == 1 {
		return s, nil
	}

	g := c.makegrid(utf8.RuneCountInString(s))
	g.FillByRow(s)
	return g.ReadByRow(), nil
}

// Decipher a message.
func (c *Cipher) Decipher(s string) (string, error) {
	if c.Rows == 1 {
		return s, nil
	}

	g := c.makegrid(utf8.RuneCountInString(s))
	g.FillByCol(s)
	return g.ReadByCol(), nil
}

// EnciphermentGrid returns the output tableau upon encipherment.
func (c *Cipher) enciphermentGrid(s string) (string, error) {
	if c.Rows == 1 {
		return s, nil
	}

	g := c.makegrid(utf8.RuneCountInString(s))
	g.FillByRow(s)
	return g.Printable(), nil
}

// DeciphermentGrid returns the output tableau upon encipherment.
func (c *Cipher) deciphermentGrid(s string) (string, error) {
	if c.Rows == 1 {
		return s, nil
	}

	g := c.makegrid(utf8.RuneCountInString(s))
	g.FillByCol(s)
	return g.Printable(), nil
}
