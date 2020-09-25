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

package dellaporta

import (
	"unicode/utf8"

	"github.com/merenbach/goldbug/internal/masc"
	"github.com/merenbach/goldbug/internal/pasc"
)

// Cipher implements a Della Porta cipher.
type Cipher struct {
	Alphabet string
	Caseless bool
	Key      string
	Strict   bool
}

// WrapString wraps a string a specified number of indices.
// WrapString will error out if the provided offset is negative.
func wrapString(s string, i int) string {
	// if we simply `return s[i:] + s[:i]`, we're operating on bytes, not runes
	rr := []rune(s)
	return string(rr[i:]) + string(rr[:i])
}

// WrapString wraps a string a specified number of indices.
// WrapString will error out if the provided offset is negative.
func owrapString(s string, i int) (string, error) {
	sLen := utf8.RuneCountInString(s)

	if sLen%2 != 0 {
		panic("owrapString sequence length must be divisible by two")
	}

	s = wrapString(s, sLen/2)

	// if we simply `return s[i:] + s[:i]`, we're operating on bytes, not runes
	sRunes := []rune(s)
	u, v := sRunes[:len(sRunes)/2], sRunes[len(sRunes)/2:]
	return wrapString(string(u), i) + wrapString(string(v), len(v)-i), nil
}

func (c *Cipher) maketableau() (*pasc.TabulaRecta, error) {
	alphabet := c.Alphabet
	if alphabet == "" {
		alphabet = pasc.Alphabet
	}

	tr, err := pasc.NewTabulaRecta(c.Alphabet, "", func(s string, i int) (*masc.Tableau, error) {
		ctAlphabet2, err := owrapString(s, i/2)
		if err != nil {
			return nil, err
		}

		t, err := masc.NewTableau(&masc.Configuration{Alphabet: alphabet}, ctAlphabet2, nil)
		if err != nil {
			return nil, err
		}
		t.Strict = c.Strict
		t.Caseless = c.Caseless
		return t, nil
	})
	if err != nil {
		return nil, err
	}

	tr.Caseless = c.Caseless
	return tr, nil
}

// Encipher a message.
func (c *Cipher) Encipher(s string) (string, error) {
	t, err := c.maketableau()
	if err != nil {
		return "", err
	}
	return t.Encipher(s, c.Key, nil)
}

// Decipher a message.
func (c *Cipher) Decipher(s string) (string, error) {
	t, err := c.maketableau()
	if err != nil {
		return "", err
	}
	return t.Decipher(s, c.Key, nil)
}

// Tableau for encipherment and decipherment.
func (c *Cipher) Tableau() (string, error) {
	t, err := c.maketableau()
	if err != nil {
		return "", err
	}
	return t.Printable()
}
