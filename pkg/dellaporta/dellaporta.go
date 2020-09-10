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

package dellaporta

import (
	"errors"
	"unicode/utf8"

	"github.com/merenbach/goldbug/internal/pasc"
	"github.com/merenbach/goldbug/internal/stringutil"
)

// Cipher implements a Della Porta cipher.
type Cipher struct {
	Alphabet string
	Key      string
	Strict   bool
}

func (c *Cipher) maketableau() (*pasc.ReciprocalTable, error) {
	alphabet := c.Alphabet
	if alphabet == "" {
		alphabet = pasc.Alphabet
	}

	ptAlphabet, ctAlphabet, keyAlphabet := alphabet, alphabet, alphabet

	ctAlphabets := make([]string, utf8.RuneCountInString(keyAlphabet))
	ctAlphabetLen := utf8.RuneCountInString(ctAlphabet)

	if ctAlphabetLen%2 != 0 {
		return nil, errors.New("Della Porta cipher alphabets must have even length")
	}

	halfCtAlphabetLen := ctAlphabetLen / 2

	for y := range ctAlphabets {
		ii := make([]int, ctAlphabetLen)
		for x := range ii {
			if x < halfCtAlphabetLen {
				// 	// v = 13 - x + x%13 + (x+y/2)%13
				// v = 13 + (x+y/2)%13 // <--- THIS IS THE MOST BASIC VERSION
				ii[x] = halfCtAlphabetLen + (x+y/2)%halfCtAlphabetLen
			} else {
				// v = (x - y/2) % 13 // <--- THIS IS THE MOST BASIC VERSION
				ii[x] = (x - y/2) % halfCtAlphabetLen
				// v2 := mod(x, 13) - y/2 + 13
				// log.Printf("v = %d and v2 = %d", v, v2)
				// 	// v = 13 - x + x%13 + (13+x-y/2)%13
			}
			// ii[x] = (13-x)%13 + x%13 + (x-y*sign(x-13)/2)%13
		}

		out, err := stringutil.Backpermute(ctAlphabet, ii)
		if err != nil {
			return nil, err
		}
		ctAlphabets[y] = string(out)
	}

	tr := pasc.ReciprocalTable{
		PtAlphabet:  ptAlphabet,
		KeyAlphabet: keyAlphabet,
		CtAlphabets: ctAlphabets,
		Strict:      c.Strict,
	}

	return &tr, nil
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
