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

// owrapString wraps two halves of a string in opposite directions, like gears turning outward.
// owrapString will panic if the provided offset is negative.
func owrapString(s string, i int) string {
	// if we simply `return s[i:] + s[:i]`, we're operating on bytes, not runes
	sRunes := []rune(s)
	if len(sRunes)%2 != 0 {
		panic("owrapString sequence length must be divisible by two")
	}
	u, v := sRunes[:len(sRunes)/2], sRunes[len(sRunes)/2:]
	return stringutil.WrapString(string(u), i) + stringutil.WrapString(string(v), len(v)-i)
}

func (c *Cipher) maketableau2() (*pasc.ReciprocalTable, error) {
	alphabet := c.Alphabet
	if alphabet == "" {
		alphabet = pasc.Alphabet
	}

	ptAlphabet, ctAlphabet, keyAlphabet := alphabet, alphabet, alphabet

	keyRunes := []rune(keyAlphabet)
	ctAlphabets := make([]string, len(keyRunes))

	if utf8.RuneCountInString(ctAlphabet)%2 != 0 {
		return nil, errors.New("Della Porta cipher alphabets must have even length")
	}

	ctRunes := []rune(ctAlphabet)

	for y := range keyRunes {
		out := make([]rune, len(ctAlphabet))
		for x := range out {
			if x < 13 {
				out[x] = ctRunes[13+(x+y/2)%13]
			} else {
				out[x] = ctRunes[(13+x-y/2)%13]
			}
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

func (c *Cipher) maketableau() (*pasc.ReciprocalTable, error) {
	alphabet := c.Alphabet
	if alphabet == "" {
		alphabet = pasc.Alphabet
	}

	ptAlphabet, ctAlphabet, keyAlphabet := alphabet, alphabet, alphabet

	keyRunes := []rune(keyAlphabet)
	ctAlphabets := make([]string, len(keyRunes))

	if utf8.RuneCountInString(ctAlphabet)%2 != 0 {
		return nil, errors.New("Della Porta cipher alphabets must have even length")
	}
	ctAlphabetWrapped := stringutil.WrapString(ctAlphabet, utf8.RuneCountInString(ctAlphabet)/2)

	for i := range keyRunes {
		ctAlphabets[i] = owrapString(ctAlphabetWrapped, i/2)
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
	t, err := c.maketableau2()
	if err != nil {
		return "", err
	}
	return t.Encipher(s, c.Key, nil)
}

// Decipher a message.
func (c *Cipher) Decipher(s string) (string, error) {
	t, err := c.maketableau2()
	if err != nil {
		return "", err
	}
	return t.Decipher(s, c.Key, nil)
}

// Tableau for encipherment and decipherment.
func (c *Cipher) tableau() (string, error) {
	t, err := c.maketableau()
	if err != nil {
		return "", err
	}
	return t.Printable()
}
