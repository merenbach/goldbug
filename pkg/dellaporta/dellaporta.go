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
	"fmt"
	"unicode/utf8"

	"github.com/merenbach/goldbug/internal/pasc"
	"github.com/merenbach/goldbug/pkg/masc"
)

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

// Cipher implements a Della Porta cipher.
type Cipher struct {
	alphabet string
	caseless bool
	key      string
	strict   bool

	*pasc.TabulaRecta
}

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

func WithAlphabet(s string) CipherOption {
	return func(c *Cipher) {
		c.alphabet = s
	}
}

func WithKey(s string) CipherOption {
	return func(c *Cipher) {
		c.key = s
	}
}

func NewCipher(opts ...CipherOption) (*Cipher, error) {
	c := &Cipher{alphabet: pasc.Alphabet}
	for _, opt := range opts {
		opt(c)
	}

	// ctAlphabet, err := Transform([]rune(c.alphabet))
	// if err != nil {
	// 	return nil, fmt.Errorf("could not transform alphabet: %w", err)
	// }

	tableau, err := pasc.NewTabulaRecta(
		pasc.WithCaseless(c.caseless),
		pasc.WithPtAlphabet(c.alphabet),
		pasc.WithKeyAlphabet(c.alphabet),
		pasc.WithKey(c.key),
		// pasc.WithCtAlphabet(string(ctAlphabet)),
		// pasc.WithStrict(c.strict),
		pasc.WithDictFunc(func(s string, i int) (*masc.Cipher, error) {
			ctAlphabet2, err := owrapString(s, i/2)
			if err != nil {
				return nil, err
			}

			params := []masc.ConfigOption{
				masc.WithAlphabet(s),
			}
			if c.caseless {
				params = append(params, masc.WithCaseless())
			}
			if c.strict {
				params = append(params, masc.WithStrict())
			}

			return masc.NewSimpleCipher(ctAlphabet2, params...)
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("could not create tableau: %w", err)
	}
	c.TabulaRecta = tableau

	return c, nil
}
