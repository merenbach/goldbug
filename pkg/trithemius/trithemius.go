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

package trithemius

import (
	"fmt"

	"github.com/merenbach/goldbug/internal/masc"
	"github.com/merenbach/goldbug/internal/pasc"
	"github.com/merenbach/goldbug/pkg/caesar"
)

// Cipher implements a Trithemius cipher.
// Cipher is effectively a Vigenère cipher with the plaintext alphabet as the key.
type Cipher struct {
	alphabet string
	caseless bool
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
		pasc.WithKey(c.alphabet),
		// pasc.WithCtAlphabet(string(ctAlphabet)),
		// pasc.WithStrict(c.strict),
		pasc.WithDictFunc(func(s string, i int) (*masc.Tableau, error) {
			params := []caesar.CipherOption{
				caesar.WithAlphabet(s),
				caesar.WithShift(i),
			}
			if c.caseless {
				params = append(params, caesar.WithCaseless())
			}
			if c.strict {
				params = append(params, caesar.WithStrict())
			}
			c2, err := caesar.NewCipher(params...)
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
