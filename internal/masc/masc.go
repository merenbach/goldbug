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
	PtAlphabet string
	CtAlphabet string

	Strict   bool
	Caseless bool

	pt2ct map[rune]rune
	ct2pt map[rune]rune
}

// New tableau.
func New(ptAlphabet string, ctAlphabet string) (*Tableau, error) {
	pt2ct, err := translation.Map(ptAlphabet, ctAlphabet, "")
	if err != nil {
		return nil, err
	}

	ct2pt, err := translation.Map(ctAlphabet, ptAlphabet, "")
	if err != nil {
		return nil, err
	}

	return &Tableau{
		pt2ct: pt2ct,
		ct2pt: ct2pt,
	}, nil
}

// EncipherRune enciphers a rune.
func (t *Tableau) EncipherRune(r rune) rune {
	// if t.Caseless {
	// 	ptAlphabet = strings.ToUpper(ptAlphabet) + strings.ToLower(ptAlphabet)
	// 	ctAlphabet = strings.ToUpper(ctAlphabet) + strings.ToLower(ctAlphabet)
	// }

	if newRune, ok := t.pt2ct[r]; ok {
		return newRune
	}

	return (-1)
}

// DecipherRune deciphers a rune.
func (t *Tableau) DecipherRune(r rune) rune {
	// if t.Caseless {
	// 	ptAlphabet = strings.ToUpper(ptAlphabet) + strings.ToLower(ptAlphabet)
	// 	ctAlphabet = strings.ToUpper(ctAlphabet) + strings.ToLower(ctAlphabet)
	// }

	if newRune, ok := t.ct2pt[r]; ok {
		return newRune
	}

	return (-1)
}

// Encipher a string.
func (t *Tableau) Encipher(s string) (string, error) {
	ptAlphabet := t.PtAlphabet
	ctAlphabet := t.CtAlphabet

	if t.Caseless {
		ptAlphabet = strings.ToUpper(ptAlphabet) + strings.ToLower(ptAlphabet)
		ctAlphabet = strings.ToUpper(ctAlphabet) + strings.ToLower(ctAlphabet)
	}

	tt := translation.Table{
		Src:    ptAlphabet,
		Dst:    ctAlphabet,
		Strict: t.Strict,
	}
	return tt.Translate(s)
}

// Decipher a string.
func (t *Tableau) Decipher(s string) (string, error) {
	ptAlphabet := t.PtAlphabet
	ctAlphabet := t.CtAlphabet

	if t.Caseless {
		ptAlphabet = strings.ToUpper(ptAlphabet) + strings.ToLower(ptAlphabet)
		ctAlphabet = strings.ToUpper(ctAlphabet) + strings.ToLower(ctAlphabet)
	}

	tt := translation.Table{
		Src:    ctAlphabet,
		Dst:    ptAlphabet,
		Strict: t.Strict,
	}
	return tt.Translate(s)
}

// Printable representation of this tableau.
func (t *Tableau) Printable() (string, error) {
	return fmt.Sprintf("PT: %s\nCT: %s", t.PtAlphabet, t.CtAlphabet), nil
}
