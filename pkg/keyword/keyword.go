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
	Caseless bool
}

func (c *Cipher) maketableau() (*masc.Tableau, error) {
	ptAlphabet := c.Alphabet
	if ptAlphabet == "" {
		ptAlphabet = masc.Alphabet
	}
	ctAlphabet := stringutil.Key(ptAlphabet, c.Keyword)

	t, err := masc.New(ptAlphabet, func(string) (string, error) {
		return ctAlphabet, nil
	})
	if err != nil {
		return nil, err
	}
	t.Strict = c.Strict
	t.Caseless = c.Caseless
	return t, nil
}

// Encipher a message.
func (c *Cipher) Encipher(s string) (string, error) {
	t, err := c.maketableau()
	if err != nil {
		log.Println("Could not calculate alphabets")
		return "", err
	}
	return t.Encipher(s)
}

// Decipher a message.
func (c *Cipher) Decipher(s string) (string, error) {
	t, err := c.maketableau()
	if err != nil {
		log.Println("Could not calculate alphabets")
		return "", err
	}
	return t.Decipher(s)
}

// Tableau for this cipher.
func (c *Cipher) Tableau() (string, error) {
	t, err := c.maketableau()
	if err != nil {
		log.Println("Could not calculate alphabets")
		return "", err
	}
	return t.Printable()
}
