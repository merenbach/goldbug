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
	"strings"
	"unicode/utf8"
)

// A Square implements a Polybius square.
type Square struct {
	Alphabet string
	Columns  int
}

func (ps *Square) String() string {
	return fmt.Sprintf("Polybius square with %d columns and alphabet %q", ps.Columns, ps.Alphabet)
}

// Printable representation of this tableau.
func (ps *Square) Printable() string {
	var b strings.Builder
	b.WriteString("   ")
	for i := 0; i < ps.Columns; i++ {
		b.WriteString(fmt.Sprintf(" %d", i+1))
	}

	b.WriteRune('\n')
	b.WriteString("  +")
	for i := 0; i < 2*ps.Columns; i++ {
		b.WriteRune('-')
	}
	alphaRunes := []rune(ps.Alphabet)
	for i := 0; i < ps.Rows(); i++ {
		b.WriteRune('\n')
		b.WriteString(fmt.Sprintf("%d | ", i+1))
		for j := 0; j < ps.Columns; j++ {
			b.WriteRune(alphaRunes[i*ps.Columns+j])
			b.WriteRune(' ')
		}
	}
	return b.String()
}

// func (ps *Square) rowFor(r rune) int {
// 	i := strings.IndexRune(ps.Alphabet, r)
// 	if i != (-1) {
// 		return i / ps.Columns
// 	}
// 	return (-1)
// }

// func (ps *Square) colFor(r rune) int {
// 	i := strings.IndexRune(ps.Alphabet, r)
// 	if i != (-1) {
// 		return i % ps.Columns
// 	}
// 	return (-1)
// }

// Rows in this Polybius square.
func (ps *Square) Rows() int {
	runeCount := utf8.RuneCountInString(ps.Alphabet)
	out := runeCount / ps.Columns
	if runeCount%ps.Columns > 0 {
		out++
	}
	return out
}

// Encipher a message.
func (ps *Square) Encipher(s string) ([]int, error) {
	pt2ct := make(map[rune]int)
	for _, r := range []rune(ps.Alphabet) {
		rIdx := strings.IndexRune(ps.Alphabet, r)
		q, m := rIdx/ps.Columns+1, rIdx%ps.Columns+1
		pt2ct[r] = 10*q + m
	}

	// TODO: substitute characters J=>I, etc.
	// TODO: option for swapped rows/cols
	// TODO: alternative side alphabets
	// TODO: nulls in alphabet
	var out []int
	for _, r := range []rune(s) {
		if o, ok := pt2ct[r]; ok {
			out = append(out, o)
		}
	}
	return out, nil
}

// Decipher a message.
func (ps *Square) Decipher(ii []int) (string, error) {
	ct2pt := make(map[int]rune)
	for _, r := range []rune(ps.Alphabet) {
		rIdx := strings.IndexRune(ps.Alphabet, r)
		q, m := rIdx/ps.Columns+1, rIdx%ps.Columns+1
		ct2pt[10*q+m] = r
	}

	var out []rune
	for _, i := range ii {
		if o, ok := ct2pt[i]; ok {
			out = append(out, o)
		}
	}
	return string(out), nil
}
