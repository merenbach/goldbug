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

package polybius

import (
	"fmt"
	"strconv"
	"strings"
	"text/tabwriter"
)

const defaultAlphabet = "ABCDEFGHIKLMNOPQRSTUVWXYZ"

// A Cipher implements a Polybius square cipher.
type Cipher struct {
	Alphabet  string
	Cols      int
	Delimiter string
	Strict    bool
}

func (c *Cipher) String() string {
	alphabet := c.Alphabet
	if alphabet == "" {
		alphabet = defaultAlphabet
	}

	var b strings.Builder

	w := tabwriter.NewWriter(&b, 0, 1, 1, ' ', 0)

	fmt.Fprintf(w, "")
	for i := 0; i < c.Cols; i++ {
		fmt.Fprintf(w, "\t%d", i+1)
	}

	alphaRunes := []rune(alphabet)
	rows := len(alphaRunes) / c.Cols
	if len(alphaRunes)%c.Cols > 0 {
		rows++
	}
	for i := 0; i < rows; i++ {
		fmt.Fprintf(w, "\n%d", i+1)
		for j := 0; j < c.Cols; j++ {
			fmt.Fprintf(w, "\t%c", alphaRunes[i*c.Cols+j])
		}
	}

	w.Flush()
	return b.String()
}

// // Makegrid creates a grid and numbers its cells.
// func (c *Cipher) makegrid(n int, cols int) grid.Grid {
// 	g := make(grid.Grid, n)
// 	for i := range g {
// 		g[i].Col = i%cols + 1
// 		g[i].Row = i/cols + 1
// 	}
// 	return g
// }

// Encipher a message.
// func (c *Cipher) Encipher(s string) (string, error) {
// 	alphabet := c.Alphabet
// 	if alphabet == "" {
// 		alphabet = defaultAlphabet
// 	}

// 	g := c.makegrid(utf8.RuneCountInString(alphabet), c.Cols)
// 	g.SortByRow()
// 	g.Fill(alphabet)

// 	m := make(map[rune]int)
// 	log.Println("len of g = ", g)
// 	for _, c := range g {
// 		m[c.Rune] = 10*c.Row + c.Col
// 	}

// 	for _, r := range []rune(s) {
// 		log.Printf("%c encodes to %d", r, m[r])
// 	}
// 	// for _, k := range c.Keys {
// 	// 	g := c.makegrid(utf8.RuneCountInString(s), utf8.RuneCountInString(k))
// 	// 	g.SortByRow()
// 	// 	g.Fill(s)

// 	// 	keyNums := lexicalKey(k, c.Myszkowski)
// 	// 	s = g.ReadCols(keyNums)
// 	// }

// 	return s, nil
// }

// // Decipher a message.
// func (c *Cipher) Decipher(s string) (string, error) {
// 	alphabet := c.Alphabet
// 	if alphabet == "" {
// 		alphabet = defaultAlphabet
// 	}

// 	g := c.makegrid(utf8.RuneCountInString(alphabet), c.Cols)
// 	g.SortByRow()
// 	g.Fill(s)

// 	// for i := len(c.Keys) - 1; i >= 0; i-- {
// 	// 	k := c.Keys[i]
// 	// 	g := c.makegrid(utf8.RuneCountInString(s), utf8.RuneCountInString(k))
// 	// 	keyNums := lexicalKey(k, c.Myszkowski)

// 	// 	g.OrderByCol(keyNums)
// 	// 	g.Fill(s)
// 	// 	g.SortByCol()
// 	// 	s = g.ReadByRow()
// 	// }

// 	return s, nil
// }

func (c *Cipher) rowcol(i int) string {
	q, m := i/c.Cols+1, i%c.Cols+1
	return strconv.Itoa(10*q + m)
}

// Encipher a message.
func (c *Cipher) Encipher(s string) (string, error) {
	alphabet := c.Alphabet
	if alphabet == "" {
		alphabet = defaultAlphabet
	}

	// create rune map
	pt2ct := make(map[rune]string)
	for i, r := range []rune(alphabet) {
		pt2ct[r] = c.rowcol(i)
	}

	// TODO: substitute characters J=>I, etc.
	// TODO: option for swapped rows/cols
	// TODO: alternative side alphabets
	// TODO: nulls in alphabet
	// TODO: strict mode
	sRunes := []rune(s)
	out := make([]string, len(sRunes))
	for i, r := range sRunes {
		if o, ok := pt2ct[r]; ok {
			out[i] = o
		}
	}

	return strings.Join(out, c.Delimiter), nil
}

// Decipher a message.
func (c *Cipher) Decipher(s string) (string, error) {
	alphabet := c.Alphabet
	if alphabet == "" {
		alphabet = defaultAlphabet
	}

	// create rune map
	ct2pt := make(map[string]rune)
	for i, r := range []rune(alphabet) {
		ct2pt[c.rowcol(i)] = r
	}

	f := strings.Split(s, c.Delimiter)
	// sRunes := []rune(s)
	// chunks := make([]string, 0)
	// for i := 0; i < len(sRunes); i += 2 {
	// 	chunks = appen
	// }

	out := make([]rune, len(f))
	for i, v := range f {
		if o, ok := ct2pt[v]; ok {
			out[i] = o
		}
	}
	return string(out), nil
}
