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
	"github.com/merenbach/goldbug/internal/masc"
	"github.com/merenbach/goldbug/internal/pasc"
	"github.com/merenbach/goldbug/pkg/affine"
)

// Cipher implements a Beaufort cipher.
// Cipher is effectively a Vigenère cipher with the ciphertext and key alphabets both mirrored (back-to-front).
type Cipher struct {
	Alphabet string
	Caseless bool
	Key      string
	Strict   bool
}

func (c *Cipher) maketableau() (*pasc.TabulaRecta, error) {
	tr, err := pasc.NewTabulaRectaFromAffineFunction(c.Alphabet, "", func(s string, i int) (*masc.Tableau, error) {
		c2 := &affine.Cipher{
			Alphabet:  s,
			Slope:     (-1),
			Intercept: i,
		}
		return c2.Tableau()
	})
	if err != nil {
		return nil, err
	}

	tr.Caseless = c.Caseless
	tr.Strict = c.Strict
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
