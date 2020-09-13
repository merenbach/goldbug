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
	"unicode"

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

	pt2ct translation.Table
	ct2pt translation.Table
}

// New tableau.
func New(ptAlphabet string, ctAlphabet string) (*Tableau, error) {
	pt2ct, err := translation.New(ptAlphabet, ctAlphabet, "")
	if err != nil {
		return nil, err
	}

	ct2pt, err := translation.New(ctAlphabet, ptAlphabet, "")
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
	if !t.Caseless {
		return t.pt2ct.Get(r, true)
	}

	r1, r2 := unicode.ToUpper(r), unicode.ToLower(r)

	// TODO: Consider putting this in the translation table methods?
	if out := t.pt2ct.Get(r1, true); out != (-1) {
		return out
	} else if out := t.pt2ct.Get(r2, true); out != (-1) {
		return out
	}

	return (-1)
}

// DecipherRune deciphers a rune.
func (t *Tableau) DecipherRune(r rune) rune {
	if !t.Caseless {
		return t.ct2pt.Get(r, true)
	}

	r1, r2 := unicode.ToUpper(r), unicode.ToLower(r)

	if out := t.ct2pt.Get(r1, true); out != (-1) {
		return out
	} else if out := t.ct2pt.Get(r2, true); out != (-1) {
		return out
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

	tt, err := translation.New(ptAlphabet, ctAlphabet, "")
	if err != nil {
		return "", err
	}

	return strings.Map(func(r rune) rune {
		return tt.Get(r, t.Strict)
	}, s), nil
}

// Decipher a string.
func (t *Tableau) Decipher(s string) (string, error) {
	ptAlphabet := t.PtAlphabet
	ctAlphabet := t.CtAlphabet

	if t.Caseless {
		ptAlphabet = strings.ToUpper(ptAlphabet) + strings.ToLower(ptAlphabet)
		ctAlphabet = strings.ToUpper(ctAlphabet) + strings.ToLower(ctAlphabet)
	}

	tt, err := translation.New(ctAlphabet, ptAlphabet, "")
	if err != nil {
		return "", err
	}

	return strings.Map(func(r rune) rune {
		return tt.Get(r, t.Strict)
	}, s), nil
}

// Printable representation of this tableau.
func (t *Tableau) Printable() (string, error) {
	return fmt.Sprintf("PT: %s\nCT: %s", t.PtAlphabet, t.CtAlphabet), nil
}
