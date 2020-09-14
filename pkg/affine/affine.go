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
	Caseless  bool
	Intercept int
	Slope     int
	Strict    bool
}

// Encipher a message.
func (c *Cipher) Encipher(s string) (string, error) {
	t, err := c.Tableau()
	if err != nil {
		log.Println("Could not calculate alphabets")
		return "", err
	}
	return t.Encipher(s)
}

// Decipher a message.
func (c *Cipher) Decipher(s string) (string, error) {
	t, err := c.Tableau()
	if err != nil {
		log.Println("Could not calculate alphabets")
		return "", err
	}
	return t.Decipher(s)
}

// EncipherRune enciphers a rune.
func (c *Cipher) EncipherRune(r rune) (rune, bool) {
	t, err := c.Tableau()
	if err != nil {
		log.Println("Could not calculate alphabets")
		return -1, false
	}
	return t.EncipherRune(r)
}

// DecipherRune deciphers a rune.
func (c *Cipher) DecipherRune(r rune) (rune, bool) {
	t, err := c.Tableau()
	if err != nil {
		log.Println("Could not calculate alphabets")
		return -1, false
	}
	return t.DecipherRune(r)
}

// Tableau for this cipher.
func (c *Cipher) Tableau() (*masc.Tableau, error) {
	t, err := masc.New(c.Alphabet, func(s string) (string, error) {
		return transform(s, c.Slope, c.Intercept)
	})
	if err != nil {
		return nil, err
	}
	t.Strict = c.Strict
	t.Caseless = c.Caseless
	return t, nil
}
