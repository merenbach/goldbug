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

	"github.com/merenbach/goldbug/internal/translation"
)

// Alphabet to use by default for substitution ciphers
const Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// A Tableau holds a translation table.
type Tableau struct {
	ptAlphabet string
	ctAlphabet string
	strict     bool
	caseless   bool

	pt2ct translation.Table
	ct2pt translation.Table
}

func NewTableau(ptAlphabet string, ctAlphabet string, strict bool, caseless bool) (*Tableau, error) {
	pt2ct, err := translation.NewTable(ptAlphabet, ctAlphabet, "")
	if err != nil {
		return nil, fmt.Errorf("could not generate pt2ct table: %w", err)
	}

	ct2pt, err := translation.NewTable(ctAlphabet, ptAlphabet, "")
	if err != nil {
		return nil, fmt.Errorf("could not generate ct2pt table: %w", err)
	}

	t := &Tableau{
		ptAlphabet,
		ctAlphabet,
		strict,
		caseless,
		pt2ct,
		ct2pt,
	}

	return t, nil
}

// EncipherRune enciphers a rune.
func (t *Tableau) EncipherRune(r rune) (rune, bool) {
	return t.pt2ct.Get(r, t.strict, t.caseless)
}

// DecipherRune deciphers a rune.
func (t *Tableau) DecipherRune(r rune) (rune, bool) {
	return t.ct2pt.Get(r, t.strict, t.caseless)
}

// Encipher a string.
func (t *Tableau) Encipher(s string) (string, error) {
	return t.pt2ct.Map(s, t.strict, t.caseless), nil
}

// Decipher a string.
func (t *Tableau) Decipher(s string) (string, error) {
	return t.ct2pt.Map(s, t.strict, t.caseless), nil
}

func (t *Tableau) String() string {
	ctAlphabet, _ := t.Encipher(t.ptAlphabet)
	return fmt.Sprintf("PT: %s\nCT: %s", t.ptAlphabet, ctAlphabet)
}
