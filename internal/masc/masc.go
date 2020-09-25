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

// A Configuration for a tableau.
type Configuration struct {
	Alphabet string

	Strict   bool
	Caseless bool
}

// A Tableau holds a translation table.
type Tableau struct {
	*Configuration

	pt2ct translation.Table
	ct2pt translation.Table
}

// NewTableau creates a new tableau.
func NewTableau(config *Configuration, ctAlphabet string, f func(string) (string, error)) (*Tableau, error) {
	if config.Alphabet == "" {
		config.Alphabet = Alphabet
	}

	if ctAlphabet == "" {
		ctAlphabet = config.Alphabet
	}

	ctAlphabetTransformed := ctAlphabet
	if f != nil {
		// Allow overrides to ctAlphabet mapping
		var err error
		ctAlphabetTransformed, err = f(ctAlphabet)
		if err != nil {
			return nil, err
		}
	}

	pt2ct, err := translation.NewTable(config.Alphabet, ctAlphabetTransformed, "")
	if err != nil {
		return nil, err
	}

	ct2pt, err := translation.NewTable(ctAlphabetTransformed, config.Alphabet, "")
	if err != nil {
		return nil, err
	}

	t := &Tableau{
		Configuration: config,
		pt2ct:         pt2ct,
		ct2pt:         ct2pt,
	}
	return t, nil
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
	return t.pt2ct.Map(s, t.Strict, t.Caseless), nil
}

// Decipher a string.
func (t *Tableau) Decipher(s string) (string, error) {
	return t.ct2pt.Map(s, t.Strict, t.Caseless), nil
}

func (t *Tableau) String() string {
	ctAlphabet := t.pt2ct.Map(t.Alphabet, t.Strict, t.Caseless)
	return fmt.Sprintf("PT: %s\nCT: %s", t.Alphabet, ctAlphabet)
}
