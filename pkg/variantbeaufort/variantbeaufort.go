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

package variantbeaufort

import (
	"github.com/merenbach/goldbug/internal/pasc"
	"github.com/merenbach/goldbug/internal/stringutil"
)

// Cipher implements a variant Beaufort cipher.
type Cipher struct {
	Alphabet string
	Key      string
	Strict   bool
}

// NewCipher creates a new variant Beaufort cipher.
func newCipher(countersign string, alphabet string) (*pasc.VigenereFamilyCipher, error) {
	if alphabet == "" {
		alphabet = pasc.Alphabet
	}
	revAlphabet := stringutil.Reverse(alphabet)
	return pasc.NewVigenereFamilyCipher(countersign, revAlphabet, revAlphabet, alphabet)
}

// Encipher a message.
func (c *Cipher) Encipher(s string) (string, error) {
	c2, err := newCipher(c.Key, c.Alphabet)
	if err != nil {
		return "", err
	}

	c2.Strict = c.Strict
	return c2.EncipherString(s), nil
}

// Decipher a message.
func (c *Cipher) Decipher(s string) (string, error) {
	c2, err := newCipher(c.Key, c.Alphabet)
	if err != nil {
		return "", err
	}

	c2.Strict = c.Strict
	return c2.DecipherString(s), nil
}

// Tableau for encipherment and decipherment.
func (c *Cipher) tableau() (string, error) {
	c2, err := newCipher(c.Key, c.Alphabet)
	if err != nil {
		return "", err
	}

	c2.Strict = c.Strict
	return c2.Printable(), nil
}
