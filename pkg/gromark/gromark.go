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

package gromark

import (
	"log"
	"sort"
	"strings"
	"unicode/utf8"

	"github.com/merenbach/goldbug/internal/grid"
	"github.com/merenbach/goldbug/internal/lfg"
	"github.com/merenbach/goldbug/internal/pasc"
	"github.com/merenbach/goldbug/internal/stringutil"
)

// make alphabetic grid based on transposition cipher
func makegrid(s string, cols []int) string {
	g := make(grid.Grid, utf8.RuneCountInString(s))

	for i := range g {
		g[i].Col = i % len(cols)
		g[i].Row = i / len(cols)
	}

	g.SortByRow()
	g.Fill(s)
	return g.ReadCols(cols)
}

// LexicalKey returns a key based on the relative lexicographic ordering of runes in a string.
func lexicalKey(s string) []int {
	out := make([]int, utf8.RuneCountInString(s))

	data := []rune(s)
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})
	dataString := string(data)

	for i, r := range []rune(s) {
		out[i] = strings.IndexRune(dataString, r) + 1
	}
	return out
}

func chainadder(m int, count int, primer []int) ([]int, error) {
	g := lfg.Additive{
		Seed:    primer,
		Taps:    []int{1, 2},
		Modulus: 10,
	}

	out := make([]int, len(primer))
	copy(out, primer)
	sliced, err := g.Slice(count)
	if err != nil {
		return nil, err
	}
	out = append(out, sliced...)
	return out, nil
}

// msglen does not need to be exact; it simply needs to meet the minimum key character needs of the cipher
func makekey(k string, msglen int) (string, error) {
	primer := make([]int, utf8.RuneCountInString(k))
	for i, r := range k {
		primer[i] = int(r - '0')
	}

	raw, err := chainadder(10, msglen, primer)
	if err != nil {
		return "", err
	}

	var b strings.Builder
	for _, r := range raw {
		b.WriteRune(rune(r + '0'))
	}

	return b.String(), nil
}

// Cipher implements a GROMARK (GROnsfeld with Mixed Alphabet and Running Key) cipher.
// Cipher key is simply the primer for a running key.
type Cipher struct {
	Alphabet string
	Key      string
	Primer   string
	Strict   bool
}

func (c *Cipher) maketableau() (*pasc.TabulaRecta, error) {
	const digits = "0123456789"

	alphabet := c.Alphabet
	if alphabet == "" {
		alphabet = pasc.Alphabet
	}

	ctAlphabetInput := stringutil.Deduplicate(c.Key + alphabet)
	cols := lexicalKey(c.Key)

	return &pasc.TabulaRecta{
		PtAlphabet:  alphabet,
		CtAlphabet:  makegrid(ctAlphabetInput, cols),
		KeyAlphabet: digits,
		Strict:      c.Strict,
	}, nil
}

// Encipher a message.
func (c *Cipher) Encipher(s string) (string, error) {
	t, err := c.maketableau()
	if err != nil {
		return "", err
	}
	log.Println(lexicalKey(c.Key))
	key, err := makekey(c.Primer, utf8.RuneCountInString(s))
	if err != nil {
		return "", err
	}
	return t.Encipher(s, key, nil)
}

// Decipher a message.
func (c *Cipher) Decipher(s string) (string, error) {
	t, err := c.maketableau()
	if err != nil {
		return "", err
	}

	key, err := makekey(c.Primer, utf8.RuneCountInString(s))
	if err != nil {
		return "", err
	}
	return t.Decipher(s, key, nil)
}

// Tableau for encipherment and decipherment.
func (c *Cipher) tableau() (string, error) {
	t, err := c.maketableau()
	if err != nil {
		return "", err
	}
	return t.Printable()
}
