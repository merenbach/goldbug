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
	"fmt"
	"sort"

	"github.com/merenbach/goldbug/internal/iterutil"
	"github.com/merenbach/goldbug/internal/sliceutil"
	"golang.org/x/exp/constraints"
)

// // LexicalKey returns a key based on the relative lexicographic ordering of runes in a string.
// func lexicalKey(s string, repeats bool) []int {
// 	out := make([]int, utf8.RuneCountInString(s))

// 	data := []rune(s)
// 	sort.Slice(data, func(i, j int) bool {
// 		return data[i] < data[j]
// 	})

// 	if repeats {
// 		data = sliceutil.Deduplicate(data)
// 	}

// 	seen := make(map[int]struct{})
// 	for i, r := range []rune(s) {

// 		var z int
// 		for i2, r2 := range data {
// 			if r2 == r {
// 				if _, ok := seen[i2]; ok {
// 					if !repeats {
// 						continue
// 					}
// 				}

// 				seen[i2] = struct{}{}
// 				z = i2
// 				break
// 			}
// 		}
// 		out[i] = z + 1
// 	}
// 	return out
// }

// Cipher implements a columnar transposition cipher.
type Cipher struct {
	*Config

	// Nulls []rune
}

func NewCipher(opts ...ConfigOption) (*Cipher, error) {
	c := NewConfig(opts...)
	return &Cipher{
		Config: c,
	}, nil
}

// Lexorder returns the relative lexical ordering of a sequence.
// This is spiritually similar to a Schwartzian transform or decorate-sort-undecorate.
func lexorder[T constraints.Ordered](xs []T) []int {
	set := sliceutil.Deduplicate(xs)
	sort.SliceStable(set, func(i int, j int) bool {
		return set[i] < set[j]
	})

	nums := make([]int, len(set))
	for i := range nums {
		nums[i] = i
	}

	// Assign cardinal positions to input characters based on sort order.
	m, _ := sliceutil.Zipmap(set, nums)

	// Map each input character to its first-seen position.
	return sliceutil.Map(xs, func(e T) int {
		return m[e]
	})
}

// TODO: add more tests to ensure that Myszkowski argument maybe always has an effect
/// Generate transposition cipher indices based on a columnar key.
func (c *Cipher) process(count int, key []int) ([]int, error) {
	// This lexical ordering transformation is technically needed only to support Myszkowski transposition.
	// We don't know the message length before here, so we want to avoid doing any argsorts until now,
	// as argsort will convert duplicate values into consecutive values.
	// key := strings.Join(c.Keys, "")
	// key := c.Keys[0]

	lexkey := lexorder(key)

	if c.myszkowski {
		xs := iterutil.Take(count, sliceutil.Cycle(lexkey))
		return sliceutil.Argsort(sliceutil.Argsort(xs)), nil
	} else {
		xs := sliceutil.Argsort(sliceutil.Argsort(lexkey))
		return iterutil.Take(count, sliceutil.Cycle(xs)), nil
	}
}

// Encipher a message.
func (c *Cipher) Encipher(xs string) (string, error) {
	// let ys: Vec<_> = xs.iter().chain(self.nulls.iter()).copied().collect();
	out := xs
	for _, key := range c.keys {
		ys := []rune(out)
		indices, err := c.process(len(ys), key)
		if err != nil {
			return "", fmt.Errorf("could not process indices: %w", err)
		}

		out2, err := sliceutil.Backpermute(ys, sliceutil.Argsort(indices))
		if err != nil {
			return "", fmt.Errorf("could not backpermute: %w", err)
		}

		out = string(out2)
	}
	return out, nil
}

// Decipher a message.
func (c *Cipher) Decipher(xs string) (string, error) {
	out := xs
	for i := len(c.keys) - 1; i >= 0; i-- {
		key := c.keys[i]

		ys := []rune(out)
		indices, err := c.process(len(ys), key)
		if err != nil {
			return "", fmt.Errorf("could not process indices: %w", err)
		}

		// TODO: this doesn't verify that the nulls are again present at the end
		out2, err := sliceutil.Backpermute(ys, sliceutil.Argsort(sliceutil.Argsort(indices)))
		if err != nil {
			return "", fmt.Errorf("could not backpermute: %w", err)
		}

		out = string(out2)
	}
	return out, nil // - len(self.nulls)
}

// Makegrid creates a grid and numbers its cells.
// func (c *Cipher) makegrid(n int, cols int) grid.Grid {
// 	g := make(grid.Grid, n)
// 	for i := range g {
// 		g[i].Col = i % cols
// 		g[i].Row = i / cols
// 	}
// 	return g
// }

// // Encipher a message.
// func (c *Cipher) Encipher(s string) (string, error) {
// 	for _, k := range c.Keys {
// 		g := c.makegrid(utf8.RuneCountInString(s), utf8.RuneCountInString(k))
// 		g.SortByRow()
// 		g.Fill(s)

// 		keyNums := lexicalKey(k, c.myszkowski)
// 		s = g.ReadCols(keyNums)
// 	}

// 	return s, nil
// }

// // Decipher a message.
// func (c *Cipher) Decipher(s string) (string, error) {
// 	for i := len(c.Keys) - 1; i >= 0; i-- {
// 		k := c.Keys[i]
// 		g := c.makegrid(utf8.RuneCountInString(s), utf8.RuneCountInString(k))
// 		keyNums := lexicalKey(k, c.myszkowski)

// 		g.OrderByCol(keyNums)
// 		g.Fill(s)
// 		g.SortByCol()
// 		s = g.ReadByRow()
// 	}

// 	return s, nil
// }

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
