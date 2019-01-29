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

package main

// Polyalphabetic substitution ciphers

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// TabulaRecta holds a tabula recta.
type TabulaRecta struct {
	ptAlphabet  string
	ctAlphabet  string
	keyAlphabet string
	ciphers     map[rune]*SimpleSubstitutionCipher
}

func (tr TabulaRecta) String() string {
	var out strings.Builder
	formatForPrinting := func(s string) string {
		spl := strings.Split(s, "")
		return strings.Join(spl, " ")
	}
	out.WriteString("    " + formatForPrinting(tr.ptAlphabet) + "\n  +")
	for range tr.ptAlphabet {
		out.WriteRune('-')
		out.WriteRune('-')
	}
	for _, r := range tr.keyAlphabet {
		c := tr.ciphers[r]
		ctAlpha := fmt.Sprintf("\n%c | %s", r, formatForPrinting(c.ctAlphabet))
		out.WriteString(ctAlpha)
	}
	return out.String()
}

// A VigenereFamilyCipher represents a cipher in the Vigenere family.
type VigenereFamilyCipher struct {
	TabulaRecta
	countersign   string
	Textautoclave bool
	Keyautoclave  bool
}

// NewVigenereFamilyCipher creates a new tabula recta suitable for use with the Vigenere family of ciphers.
// NOTE: we roll the countersign into the tabula recta so it has all the data it needs
// to decode/encode a string reusably, for parallelism with the monoalphabetic ciphers.
func NewVigenereFamilyCipher(countersign, ptAlphabet, ctAlphabet, keyAlphabet string) *VigenereFamilyCipher {
	tr := TabulaRecta{
		ptAlphabet:  ptAlphabet,
		ctAlphabet:  ctAlphabet,
		keyAlphabet: keyAlphabet,
		ciphers:     make(map[rune]*SimpleSubstitutionCipher),
	}
	// this cast is necessary to ensure that the index increases without gaps
	for i, r := range []rune(keyAlphabet) {
		ctAlphabet3 := wrapString(ctAlphabet, i)
		tr.ciphers[r] = NewSimpleSubstitutionCipher(ptAlphabet, ctAlphabet3)
	}
	return &VigenereFamilyCipher{
		TabulaRecta: tr,
		countersign: countersign,
	}
}

// NewDellaPortaReciprocalTable creates a new tabula recta suitable for use with the Della Porta cipher.
func NewDellaPortaReciprocalTable(countersign, ptAlphabet, ctAlphabet, keyAlphabet string) *VigenereFamilyCipher {
	tr := TabulaRecta{
		ptAlphabet:  ptAlphabet,
		ctAlphabet:  ctAlphabet,
		keyAlphabet: keyAlphabet,
		ciphers:     make(map[rune]*SimpleSubstitutionCipher),
	}
	if utf8.RuneCountInString(ctAlphabet)%2 != 0 {
		panic("Della Porta cipher alphabets must have even length")
	}
	ctAlphabet2 := wrapString(ctAlphabet, utf8.RuneCountInString(ctAlphabet)/2)
	// this cast is necessary to ensure that the index increases without gaps
	for i, r := range []rune(keyAlphabet) {
		ctAlphabet3 := owrapString(ctAlphabet2, i/2)
		tr.ciphers[r] = NewSimpleSubstitutionCipher(ptAlphabet, ctAlphabet3)
	}
	return &VigenereFamilyCipher{
		TabulaRecta: tr,
		countersign: countersign,
	}
}

// Encipher a message from plaintext to ciphertext.
func (c VigenereFamilyCipher) Encipher(s string, strict bool) string {
	keyRunes := []rune(c.countersign)
	var transcodedCharCount = 0
	return strings.Map(func(r rune) rune {
		k := keyRunes[transcodedCharCount%len(keyRunes)]
		cipher := c.ciphers[k]
		o, ok := cipher.EncipherRune(r, strict)
		if ok {
			transcodedCharCount++
			if c.Textautoclave {
				keyRunes = append(keyRunes, r)
			} else if c.Keyautoclave {
				keyRunes = append(keyRunes, o)
			}
		}
		return o
	}, s)
}

// Decipher a message from ciphertext to plaintext.
func (c VigenereFamilyCipher) Decipher(s string, strict bool) string {
	keyRunes := []rune(c.countersign)
	var transcodedCharCount = 0
	return strings.Map(func(r rune) rune {
		k := keyRunes[transcodedCharCount%len(keyRunes)]
		cipher := c.ciphers[k]
		o, ok := cipher.DecipherRune(r, strict)
		if ok {
			transcodedCharCount++
			if c.Textautoclave {
				keyRunes = append(keyRunes, o)
			} else if c.Keyautoclave {
				keyRunes = append(keyRunes, r)
			}
		}
		return o
	}, s)
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
	return wrapString(string(u), i) + wrapString(string(v), len(v)-i)
}

// NewVigenereCipher creates a new Vigenere cipher.
func NewVigenereCipher(countersign, alphabet string) *VigenereFamilyCipher {
	return NewVigenereFamilyCipher(countersign, alphabet, alphabet, alphabet)
}

// NewVigenereTextAutoclaveCipher creates a new Vigenere (text autoclave) cipher.
func NewVigenereTextAutoclaveCipher(countersign, alphabet string) *VigenereFamilyCipher {
	c := NewVigenereCipher(countersign, alphabet)
	c.Textautoclave = true
	return c
}

// NewVigenereKeyAutoclaveCipher creates a new Vigenere (key autoclave) cipher.
func NewVigenereKeyAutoclaveCipher(countersign, alphabet string) *VigenereFamilyCipher {
	c := NewVigenereCipher(countersign, alphabet)
	c.Keyautoclave = true
	return c
}

// NewBeaufortCipher creates a new Beaufort cipher.
func NewBeaufortCipher(countersign, alphabet string) *VigenereFamilyCipher {
	revAlphabet := reverseString(alphabet)
	return NewVigenereFamilyCipher(countersign, alphabet, revAlphabet, revAlphabet)
}

// NewGronsfeldCipher creates a new Gronsfeld cipher.
func NewGronsfeldCipher(countersign, alphabet string) *VigenereFamilyCipher {
	return NewVigenereFamilyCipher(countersign, alphabet, alphabet, "0123456789")
}

// NewVariantBeaufortCipher creates a new Vigenere cipher.
func NewVariantBeaufortCipher(countersign, alphabet string) *VigenereFamilyCipher {
	revAlphabet := reverseString(alphabet)
	return NewVigenereFamilyCipher(countersign, revAlphabet, revAlphabet, alphabet)
}

// NewTrithemiusCipher creates a new Trithemius cipher.
// NewTrithemiusCipher considers this simply the Vigenere cipher with the countersign equal to the alphabet.
func NewTrithemiusCipher(alphabet string) *VigenereFamilyCipher {
	return NewVigenereCipher(alphabet, alphabet)
}

// NewDellaPortaCipher creates a new DellaPorta cipher.
func NewDellaPortaCipher(countersign, alphabet string) *VigenereFamilyCipher {
	return NewDellaPortaReciprocalTable(countersign, alphabet, alphabet, alphabet)
}
