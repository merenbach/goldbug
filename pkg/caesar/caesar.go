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

package caesar

import (
	"fmt"

	"github.com/merenbach/goldbug/internal/masc"
	"github.com/merenbach/goldbug/internal/translation"
)

// adapted from: https://www.sohamkamani.com/golang/options-pattern/

type CipherOption func(*Cipher)

func WithStrict() CipherOption {
	return func(c *Cipher) {
		c.strict = true
	}
}

func WithCaseless() CipherOption {
	return func(c *Cipher) {
		c.caseless = true
	}
}

func WithAlphabet(alphabet string) CipherOption {
	return func(c *Cipher) {
		c.alphabet = alphabet
	}
}

func WithShift(shift int) CipherOption {
	return func(c *Cipher) {
		c.shift = shift
	}
}

func NewCipher(opts ...CipherOption) (*Cipher, error) {
	c := &Cipher{alphabet: masc.Alphabet}
	for _, opt := range opts {
		opt(c)
	}

	ctAlphabet, err := Transform([]rune(c.alphabet), c.shift)
	if err != nil {
		return nil, fmt.Errorf("could not transform alphabet: %w", err)
	}

	pt2ct, err := translation.NewTable(c.alphabet, string(ctAlphabet), "")
	if err != nil {
		return nil, fmt.Errorf("could not create pt2ct table: %w", err)
	}

	ct2pt, err := translation.NewTable(string(ctAlphabet), c.alphabet, "")
	if err != nil {
		return nil, fmt.Errorf("could not create ct2pt table: %w", err)
	}

	c.pt2ct = pt2ct
	c.ct2pt = ct2pt

	return c, nil
}

// Cipher implements a Caesar cipher.
type Cipher struct {
	alphabet string
	caseless bool
	strict   bool
	shift    int

	pt2ct translation.Table
	ct2pt translation.Table
}

// Encipher a message.
func (c *Cipher) Encipher(s string) (string, error) {
	return c.pt2ct.Map(s, c.strict, c.caseless), nil
}

// Decipher a message.
func (c *Cipher) Decipher(s string) (string, error) {
	return c.ct2pt.Map(s, c.strict, c.caseless), nil
}

// Tableau for encipherment and decipherment.
// func (c *Cipher) Tableau() (*masc.Tableau, error) {
// 	return c.maketableau().Tableau()
// }

func (c *Cipher) Tableau() string {
	ctAlphabet, _ := c.Encipher(c.alphabet)
	return fmt.Sprintf("PT: %s\nCT: %s", c.alphabet, ctAlphabet)
}
