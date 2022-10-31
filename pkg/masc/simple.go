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

// A SimpleCipher implements a simple cipher.
type SimpleCipher struct {
	*Config

	pt2ct translation.Table
	ct2pt translation.Table
}

// NewSimpleCipher creates and returns a new simple cipher.
func NewSimpleCipher(ctAlphabet string, opts ...ConfigOption) (*SimpleCipher, error) {
	c := NewConfig(opts...)

	pt2ct, err := translation.NewTable(c.alphabet, ctAlphabet, "")
	if err != nil {
		return nil, fmt.Errorf("could not generate pt2ct table: %w", err)
	}

	ct2pt, err := translation.NewTable(ctAlphabet, c.alphabet, "")
	if err != nil {
		return nil, fmt.Errorf("could not generate ct2pt table: %w", err)
	}

	return &SimpleCipher{
		c,
		pt2ct,
		ct2pt,
	}, nil
}

// EncipherRune enciphers a rune.
func (c *SimpleCipher) EncipherRune(r rune) (rune, bool) {
	return c.pt2ct.Get(r, c.strict, c.caseless)
}

// DecipherRune deciphers a rune.
func (c *SimpleCipher) DecipherRune(r rune) (rune, bool) {
	return c.ct2pt.Get(r, c.strict, c.caseless)
}

// Encipher a string.
func (c *SimpleCipher) Encipher(s string) (string, error) {
	return c.pt2ct.Map(s, c.strict, c.caseless), nil
}

// Decipher a string.
func (c *SimpleCipher) Decipher(s string) (string, error) {
	return c.ct2pt.Map(s, c.strict, c.caseless), nil
}

func (c *SimpleCipher) String() string {
	ctAlphabet, _ := c.Encipher(c.alphabet)
	return fmt.Sprintf("PT: %s\nCT: %s", c.alphabet, ctAlphabet)
}
