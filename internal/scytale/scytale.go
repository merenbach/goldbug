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

package scytale

import (
	"unicode/utf8"

	"github.com/merenbach/goldbug/internal/grid"
)

// A Cipher implements the scytale (or skytale) cipher.
type Cipher struct {
	Turns int
}

// Makegrid creates a grid and numbers its cells.
func (c *Cipher) makegrid(n int) grid.Grid {
	g := make(grid.Grid, n)
	for i := range g {
		g[i].Col = i % c.Turns
		g[i].Row = i / c.Turns
	}
	return g
}

// Encipher a message.
func (c *Cipher) Encipher(s string) string {
	if c.Turns == 1 {
		return s
	}

	g := c.makegrid(utf8.RuneCountInString(s))
	g.FillByCol(s)
	return g.ReadByCol()
}

// Decipher a message.
func (c *Cipher) Decipher(s string) string {
	if c.Turns == 1 {
		return s
	}

	g := c.makegrid(utf8.RuneCountInString(s))
	g.FillByRow(s)
	return g.ReadByRow()
}

// EnciphermentGrid returns the output tableau upon encipherment.
func (c *Cipher) EnciphermentGrid(s string) string {
	if c.Turns == 1 {
		return s
	}

	g := c.makegrid(utf8.RuneCountInString(s))
	g.FillByCol(s)
	return g.Printable()
}

// DeciphermentGrid returns the output tableau upon encipherment.
func (c *Cipher) DeciphermentGrid(s string) string {
	if c.Turns == 1 {
		return s
	}

	g := c.makegrid(utf8.RuneCountInString(s))
	g.FillByRow(s)
	return g.Printable()
}
