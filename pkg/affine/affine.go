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

package affine

import (
	"log"

	"github.com/merenbach/goldbug/internal/masc"
)

// Cipher implements an affine cipher.
type Cipher struct {
	Alphabet  string
	Intercept int
	Slope     int
	Strict    bool
}

// newCipher creates a new affine cipher.
func newCipher(ptAlphabet string, a int, b int) (*masc.MonoalphabeticSubstitutionCipher, error) {
	if ptAlphabet == "" {
		ptAlphabet = masc.Alphabet
	}

	ctAlphabet, err := affineTransform(ptAlphabet, a, b)
	if err != nil {
		return nil, err
	}

	return masc.New(ptAlphabet, ctAlphabet)
}

// Encipher a message.
func (c *Cipher) Encipher(s string) (string, error) {
	c2, err := newCipher(c.Alphabet, c.Slope, c.Intercept)
	if err != nil {
		log.Println("Couldn't create cipher")
		return "", err
	}
	c2.Strict = c.Strict
	return c2.EncipherString(s)
}

// Decipher a message.
func (c *Cipher) Decipher(s string) (string, error) {
	c2, err := newCipher(c.Alphabet, c.Slope, c.Intercept)
	if err != nil {
		log.Println("Couldn't create cipher")
		return "", err
	}
	c2.Strict = c.Strict
	return c2.DecipherString(s)
}

// Tableau for encipherment and decipherment.
func (c *Cipher) Tableau() (string, error) {
	c2, err := newCipher(c.Alphabet, c.Slope, c.Intercept)
	if err != nil {
		log.Println("Couldn't create cipher")
		return "", err
	}
	return c2.Printable(), nil
}
