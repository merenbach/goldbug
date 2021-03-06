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

func (c *Cipher) maketableau() (*masc.Tableau, error) {
	ptAlphabet := c.Alphabet
	if ptAlphabet == "" {
		ptAlphabet = masc.Alphabet
	}
	ctAlphabet, err := transform(ptAlphabet, c.Slope, c.Intercept)
	if err != nil {
		return nil, err
	}

	return &masc.Tableau{
		PtAlphabet: ptAlphabet,
		CtAlphabet: ctAlphabet,
		Strict:     c.Strict,
	}, nil
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
