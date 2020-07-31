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

package keyword

import (
	"log"

	"github.com/merenbach/goldbug/internal/masc"
	"github.com/merenbach/goldbug/internal/stringutil"
)

// Cipher implements a keyword cipher.
type Cipher struct {
	Alphabet string
	Keyword  string
	Strict   bool
}

// NewCipher creates a new keyword cipher.
// NewCipher currently doesn't return any errors, but returns the error argument for shape consistency.
func newCipher(ptAlphabet string, keyword string) (*masc.MonoalphabeticSubstitutionCipher, error) {
	if ptAlphabet == "" {
		ptAlphabet = masc.Alphabet
	}

	ctAlphabet := stringutil.Deduplicate(keyword + ptAlphabet)
	return masc.New(ptAlphabet, ctAlphabet)
}

// Encipher a message.
func (c *Cipher) Encipher(s string) (string, error) {
	c2, err := newCipher(c.Alphabet, c.Keyword)
	if err != nil {
		log.Println("Couldn't create cipher")
		return "", err
	}
	c2.Strict = c.Strict
	return c2.EncipherString(s)
}

// Decipher a message.
func (c *Cipher) Decipher(s string) (string, error) {
	c2, err := newCipher(c.Alphabet, c.Keyword)
	if err != nil {
		log.Println("Couldn't create cipher")
		return "", err
	}
	c2.Strict = c.Strict
	return c2.DecipherString(s)
}

// Tableau for encipherment and decipherment.
func (c *Cipher) tableau() (string, error) {
	c2, err := newCipher(c.Alphabet, c.Keyword)
	if err != nil {
		log.Println("Couldn't create cipher")
		return "", err
	}
	return c2.Printable(), nil
}
