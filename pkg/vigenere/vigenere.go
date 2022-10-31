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

package vigenere

import (
	"fmt"

	"github.com/merenbach/goldbug/internal/masc"
	"github.com/merenbach/goldbug/internal/pasc"
	"github.com/merenbach/goldbug/pkg/masc2"
)

// An AutokeyOption determines the autokey setting for the cipher.
type autokeyOption uint8

const (
	// NoAutokey signifies not to autokey the cipher.
	NoAutokey autokeyOption = iota

	// TextAutokey denotes a text-autokey mechanism.
	TextAutokey

	// KeyAutokey denotes a key-autokey mechanism.
	KeyAutokey
)

// Cipher implements a Vigenère cipher.
type Cipher struct {
	autokey  autokeyOption
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

func WithAutokey(v autokeyOption) CipherOption {
	return func(c *Cipher) {
		c.autokey = v
	}
}

func NewCipher(opts ...CipherOption) (*Cipher, error) {
	c := &Cipher{alphabet: masc.Alphabet}
	for _, opt := range opts {
		opt(c)
	}

	// ctAlphabet, err := Transform([]rune(c.alphabet))
	// if err != nil {
	// 	return nil, fmt.Errorf("could not transform alphabet: %w", err)
	// }

	params := []pasc.TabulaRectaOption{
		pasc.WithCaseless(c.caseless),
		pasc.WithPtAlphabet(c.alphabet),
		pasc.WithKeyAlphabet(c.alphabet),
		pasc.WithKey(c.key),
		// pasc.WithCtAlphabet(string(ctAlphabet)),
		// pasc.WithStrict(c.strict),
		pasc.WithDictFunc(func(s string, i int) (*masc2.Cipher, error) {
			params2 := []masc2.ConfigOption{
				masc2.WithAlphabet(s),
			}
			if c.caseless {
				params2 = append(params2, masc2.WithCaseless())
			}
			if c.strict {
				params2 = append(params2, masc2.WithStrict())
			}
			return masc2.NewCaesarCipher(i, params2...)
		}),
	}
	params = append(params, pasc.WithAutokeyFunc(func(a rune, b rune, keystream *[]rune) {
		switch c.autokey {
		case TextAutokey:
			*keystream = append(*keystream, a)
		case KeyAutokey:
			*keystream = append(*keystream, b)
		case NoAutokey:
			// do nothing
		}
	}))

	tableau, err := pasc.NewTabulaRecta(params...)
	if err != nil {
		return nil, fmt.Errorf("could not create tableau: %w", err)
	}
	c.TabulaRecta = tableau

	return c, nil
}
