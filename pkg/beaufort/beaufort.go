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

package beaufort

import (
	"fmt"

	"github.com/merenbach/goldbug/internal/masc"
	"github.com/merenbach/goldbug/internal/pasc"
	"github.com/merenbach/goldbug/pkg/affine"
)

// Cipher implements a Beaufort cipher.
// Cipher is effectively a Vigenère cipher with the ciphertext and key alphabets both mirrored (back-to-front).
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
	c := &Cipher{alphabet: masc.Alphabet}
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
		pasc.WithDictFunc(func(s string, i int) (*masc.Tableau, error) {
			params := []affine.CipherOption{
				affine.WithAlphabet(s),
				affine.WithSlope(-1),
				affine.WithIntercept(i),
			}
			if c.caseless {
				params = append(params, affine.WithCaseless())
			}
			if c.strict {
				params = append(params, affine.WithStrict())
			}
			c2, err := affine.NewCipher(params...)
			if err != nil {
				return nil, fmt.Errorf("could not create cipher: %w", err)
			}

			return c2.Tableau, nil
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("could not create tableau: %w", err)
	}
	c.TabulaRecta = tableau

	return c, nil
}

// // Encipher a message.
// func (c *Cipher) Encipher(s string) (string, error) {
// 	t, err := c.maketableau()
// 	if err != nil {
// 		return "", err
// 	}
// 	return t.Encipher(s, c.Key, nil)
// }

// // Decipher a message.
// func (c *Cipher) Decipher(s string) (string, error) {
// 	t, err := c.maketableau()
// 	if err != nil {
// 		return "", err
// 	}
// 	return t.Decipher(s, c.Key, nil)
// }

// // Tableau for encipherment and decipherment.
// func (c *Cipher) Tableau() (string, error) {
// 	t, err := c.maketableau()
// 	if err != nil {
// 		return "", err
// 	}
// 	return t.Printable()
// }
