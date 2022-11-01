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

package transposition

import (
	"sort"
	"unicode/utf8"

	"github.com/merenbach/goldbug/internal/grid"
	"github.com/merenbach/goldbug/internal/sliceutil"
)

// LexicalKey returns a key based on the relative lexicographic ordering of runes in a string.
func lexicalKey(s string, repeats bool) []int {
	out := make([]int, utf8.RuneCountInString(s))

	data := []rune(s)
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})

	if repeats {
		data = sliceutil.Deduplicate(data)
	}

	seen := make(map[int]struct{})
	for i, r := range []rune(s) {

		var z int
		for i2, r2 := range data {
			if r2 == r {
				if _, ok := seen[i2]; ok {
					if !repeats {
						continue
					}
				}

				seen[i2] = struct{}{}
				z = i2
				break
			}
		}
		out[i] = z + 1
	}
	return out
}

// Cipher implements a columnar transposition cipher.
type Cipher struct {
	*Config

	Keys []string
}

func NewCipher(keys []string, opts ...ConfigOption) *Cipher {
	c := NewConfig(opts...)
	return &Cipher{
		Config: c,
		Keys:   keys,
	}
}

// Makegrid creates a grid and numbers its cells.
func (c *Cipher) makegrid(n int, cols int) grid.Grid {
	g := make(grid.Grid, n)
	for i := range g {
		g[i].Col = i % cols
		g[i].Row = i / cols
	}
	return g
}

// Encipher a message.
func (c *Cipher) Encipher(s string) (string, error) {
	for _, k := range c.Keys {
		g := c.makegrid(utf8.RuneCountInString(s), utf8.RuneCountInString(k))
		g.SortByRow()
		g.Fill(s)

		keyNums := lexicalKey(k, c.myszkowski)
		s = g.ReadCols(keyNums)
	}

	return s, nil
}

// Decipher a message.
func (c *Cipher) Decipher(s string) (string, error) {
	for i := len(c.Keys) - 1; i >= 0; i-- {
		k := c.Keys[i]
		g := c.makegrid(utf8.RuneCountInString(s), utf8.RuneCountInString(k))
		keyNums := lexicalKey(k, c.myszkowski)

		g.OrderByCol(keyNums)
		g.Fill(s)
		g.SortByCol()
		s = g.ReadByRow()
	}

	return s, nil
}

// // EnciphermentGrid returns the output tableau upon encipherment.
// func (c *Cipher) enciphermentGrid(s string) (string, error) {
// 	g := c.makegrid(utf8.RuneCountInString(s))
// 	g.SortByCol()
// 	g.Fill(s)
// 	return g.Printable(), nil
// }

// // DeciphermentGrid returns the output tableau upon encipherment.
// func (c *Cipher) deciphermentGrid(s string) (string, error) {
// 	g := c.makegrid(utf8.RuneCountInString(s))
// 	g.SortByRow()
// 	g.Fill(s)
// 	return g.Printable(), nil
// }
