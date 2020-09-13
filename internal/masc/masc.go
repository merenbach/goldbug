// Copyright 2018 Andrew Merenbach
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

package masc

import (
	"fmt"
	"strings"

	"github.com/merenbach/goldbug/internal/translation"
)

// Alphabet to use by default for substitution ciphers
const Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// A Tableau holds a translation table.
type Tableau struct {
	Alphabet string

	Strict   bool
	Caseless bool

	pt2ct translation.Table
	ct2pt translation.Table
}

// New tableau.
func New(ptAlphabet string, f func(string) (string, error)) (*Tableau, error) {
	if ptAlphabet == "" {
		ptAlphabet = Alphabet
	}

	ctAlphabet, err := f(ptAlphabet)
	if err != nil {
		return nil, err
	}

	pt2ct, err := translation.New(ptAlphabet, ctAlphabet, "")
	if err != nil {
		return nil, err
	}

	ct2pt, err := translation.New(ctAlphabet, ptAlphabet, "")
	if err != nil {
		return nil, err
	}

	return &Tableau{
		Alphabet: ptAlphabet,
		pt2ct:    pt2ct,
		ct2pt:    ct2pt,
	}, nil
}

// EncipherRune enciphers a rune.
func (t *Tableau) EncipherRune(r rune) (rune, bool) {
	return t.pt2ct.Get(r, t.Strict, t.Caseless)
}

// DecipherRune deciphers a rune.
func (t *Tableau) DecipherRune(r rune) (rune, bool) {
	return t.ct2pt.Get(r, t.Strict, t.Caseless)
}

// Encipher a string.
func (t *Tableau) Encipher(s string) (string, error) {
	return strings.Map(func(r rune) rune {
		o, _ := t.pt2ct.Get(r, t.Strict, t.Caseless)
		return o
	}, s), nil
}

// Decipher a string.
func (t *Tableau) Decipher(s string) (string, error) {
	return strings.Map(func(r rune) rune {
		o, _ := t.ct2pt.Get(r, t.Strict, t.Caseless)
		return o
	}, s), nil
}

// Printable representation of this tableau.
func (t *Tableau) Printable() (string, error) {
	ctAlphabet := strings.Map(func(r rune) rune {
		o, _ := t.pt2ct.Get(r, t.Strict, t.Caseless)
		return o
	}, t.Alphabet)
	return fmt.Sprintf("PT: %s\nCT: %s", t.Alphabet, ctAlphabet), nil
}
