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
	"github.com/merenbach/goldbug/internal/pasc"
	"github.com/merenbach/goldbug/pkg/vigenere"
)

// Cipher implements a Trithemius cipher.
type Cipher struct {
	Alphabet string
	// Caseless bool
	Strict bool
}

// Encipher a message.
func (c *Cipher) Encipher(s string) (string, error) {
	alphabet := c.Alphabet
	if alphabet == "" {
		alphabet = pasc.Alphabet
	}

	c2 := vigenere.Cipher{
		Alphabet: alphabet,
		// Caseless: c.Caseless,
		Key:    alphabet,
		Strict: c.Strict,
	}
	return c2.Encipher(s)
}

// Decipher a message.
func (c *Cipher) Decipher(s string) (string, error) {
	alphabet := c.Alphabet
	if alphabet == "" {
		alphabet = pasc.Alphabet
	}

	c2 := vigenere.Cipher{
		Alphabet: alphabet,
		Key:      alphabet,
		Strict:   c.Strict,
	}
	return c2.Decipher(s)
}

// Tableau for encipherment and decipherment.
func (c *Cipher) Tableau() (string, error) {
	alphabet := c.Alphabet
	if alphabet == "" {
		alphabet = pasc.Alphabet
	}

	c2 := vigenere.Cipher{
		Alphabet: alphabet,
		Key:      alphabet,
		Strict:   c.Strict,
	}
	return c2.Tableau()
}
