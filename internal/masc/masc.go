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
	transform  func(string) (string, error)

	pt2ct translation.Table
	ct2pt translation.Table
}

// adapted from: https://www.sohamkamani.com/golang/options-pattern/

type TableauOption func(*Tableau)

func WithStrict(b bool) TableauOption {
	return func(c *Tableau) {
		c.strict = b
	}
}

func WithCaseless(b bool) TableauOption {
	return func(c *Tableau) {
		c.caseless = b
	}
}

func WithPtAlphabet(s string) TableauOption {
	return func(c *Tableau) {
		if s != "" {
			c.ptAlphabet = s
		}
	}
}

func WithCtAlphabet(s string) TableauOption {
	return func(c *Tableau) {
		if s != "" {
			c.ctAlphabet = s
		}
	}
}

func WithTransform(f func(string) (string, error)) TableauOption {
	return func(c *Tableau) {
		c.transform = f
	}
}

func NewTableau(opts ...TableauOption) (*Tableau, error) {
	t := &Tableau{
		ptAlphabet: Alphabet,
		ctAlphabet: Alphabet,
	}
	for _, opt := range opts {
		opt(t)
	}

	ctAlphabetTransformed := t.ctAlphabet
	if t.transform != nil {
		ctAlphabet2, err := t.transform(t.ctAlphabet)
		if err != nil {
			return nil, fmt.Errorf("could not transform alphabet: %w", err)
		}
		ctAlphabetTransformed = ctAlphabet2
	}

	pt2ct, err := translation.NewTable(t.ptAlphabet, ctAlphabetTransformed, "")
	if err != nil {
		return nil, fmt.Errorf("could not generate pt2ct table: %w", err)
	}

	ct2pt, err := translation.NewTable(ctAlphabetTransformed, t.ptAlphabet, "")
	if err != nil {
		return nil, fmt.Errorf("could not generate ct2pt table: %w", err)
	}

	t.pt2ct = pt2ct
	t.ct2pt = ct2pt

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
