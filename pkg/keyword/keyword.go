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
	"fmt"
	"log"

	"github.com/merenbach/goldbug/internal/masc"
	"github.com/merenbach/goldbug/internal/stringutil"
	"github.com/merenbach/goldbug/internal/translation"
)

// Cipher implements a keyword cipher.
type Cipher struct {
	Alphabet string
	Keyword  string
	Strict   bool
}

func (c *Cipher) alphabets() (string, string, error) {
	ptAlphabet := c.Alphabet
	if ptAlphabet == "" {
		ptAlphabet = masc.Alphabet
	}
	ctAlphabet := stringutil.Deduplicate(c.Keyword + ptAlphabet)
	return ptAlphabet, ctAlphabet, nil
}

// Encipher a message.
func (c *Cipher) Encipher(s string) (string, error) {
	ptAlphabet, ctAlphabet, err := c.alphabets()
	if err != nil {
		log.Println("Could not calculate alphabets")
		return "", err
	}

	table := translation.Table{
		Src:    ptAlphabet,
		Dst:    ctAlphabet,
		Strict: c.Strict,
	}
	return table.Translate(s)
}

// Decipher a message.
func (c *Cipher) Decipher(s string) (string, error) {
	ptAlphabet, ctAlphabet, err := c.alphabets()
	if err != nil {
		log.Println("Could not calculate alphabets")
		return "", err
	}

	table := translation.Table{
		Src:    ctAlphabet,
		Dst:    ptAlphabet,
		Strict: c.Strict,
	}
	return table.Translate(s)
}

// Tableau for this cipher.
func (c *Cipher) Tableau() (string, error) {
	ptAlphabet, ctAlphabet, err := c.alphabets()
	if err != nil {
		log.Println("Could not calculate alphabets")
		return "", err
	}
	return fmt.Sprintf("PT: %s\nCT: %s", ptAlphabet, ctAlphabet), nil
}
