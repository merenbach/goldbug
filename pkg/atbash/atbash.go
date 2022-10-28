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

package atbash

import (
	"fmt"

	"github.com/merenbach/goldbug/internal/masc"
	"github.com/merenbach/goldbug/internal/translation"
	"github.com/merenbach/goldbug/pkg/affine"
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

func NewCipher(opts ...CipherOption) (*Cipher, err) {
	c := &Cipher{alphabet: masc.Alphabet}
	for _, opt := range opts {
		opt(c)
	}

	ctAlphabet, err := affine.Transform(c.Alphabet, (-1), (-1))
	if err != nil {
		return nil, fmt.Errorf("could not transform alphabet: %w", err)
	}

	c.pt2ct, err := translation.NewTable(c.Alphabet, ctAlphabet, "")
	if err != nil {
		return nil, fmt.Errorf("could not create pt2ct table: %w", err)
	}

	c.ct2pt,err := translation.NewTable(ctAlphabet, c.Alphabet, "")
	if err != nil {
		return nil, fmt.Errorf("could not create ct2pt table: %w", err)
	}

	return c, nil
}

// Cipher implements an Atbash cipher.
type Cipher struct {
	alphabet string
	caseless bool
	strict   bool

	pt2ct translation.Table
	ct2pt translation.Table
}

// Encipher a message.
func (c *Cipher) Encipher(s string) (string, error) {
	return c.pt2ct.Encipher(s)
}

// Decipher a message.
func (c *Cipher) Decipher(s string) (string, error) {
	return c.ct2pt().Decipher(s)
}

// Tableau for encipherment and decipherment.
// func (c *Cipher) Tableau() (*masc.Tableau, error) {
// 	return c.maketableau().Tableau()
// }
