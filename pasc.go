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
	ciphers     map[rune]SimpleSubstitutionCipher
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

// MakeVigenereFamilyCipher creates a new tabula recta suitable for use with the Vigenere family of ciphers.
// NOTE: we roll the countersign into the tabula recta so it has all the data it needs
// to decode/encode a string reusably, for parallelism with the monoalphabetic ciphers.
func MakeVigenereFamilyCipher(countersign, ptAlphabet, ctAlphabet, keyAlphabet string) VigenereFamilyCipher {
	tr := TabulaRecta{
		ptAlphabet:  ptAlphabet,
		ctAlphabet:  ctAlphabet,
		keyAlphabet: keyAlphabet,
		ciphers:     make(map[rune]SimpleSubstitutionCipher),
	}
	// this cast is necessary to ensure that the index increases without gaps
	for i, r := range []rune(keyAlphabet) {
		ctAlphabet3 := wrapString(ctAlphabet, i)
		tr.ciphers[r] = makeSimpleSubstitutionCipher(ptAlphabet, ctAlphabet3)
	}
	return VigenereFamilyCipher{
		TabulaRecta: tr,
		countersign: countersign,
	}
}

// MakeDellaPortaReciprocalTable creates a new tabula recta suitable for use with the Della Porta cipher.
func MakeDellaPortaReciprocalTable(countersign, ptAlphabet, ctAlphabet, keyAlphabet string) VigenereFamilyCipher {
	tr := TabulaRecta{
		ptAlphabet:  ptAlphabet,
		ctAlphabet:  ctAlphabet,
		keyAlphabet: keyAlphabet,
		ciphers:     make(map[rune]SimpleSubstitutionCipher),
	}
	if utf8.RuneCountInString(ctAlphabet)%2 != 0 {
		panic("Della Porta cipher alphabets must have even length")
	}
	ctAlphabet2 := wrapString(ctAlphabet, utf8.RuneCountInString(ctAlphabet)/2)
	// this cast is necessary to ensure that the index increases without gaps
	for i, r := range []rune(keyAlphabet) {
		ctAlphabet3 := owrapString(ctAlphabet2, i/2)
		tr.ciphers[r] = makeSimpleSubstitutionCipher(ptAlphabet, ctAlphabet3)
	}
	return VigenereFamilyCipher{
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
		if o := cipher.encipherRune(r); o != (-1) {
			transcodedCharCount++
			if c.Textautoclave {
				keyRunes = append(keyRunes, r)
			} else if c.Keyautoclave {
				keyRunes = append(keyRunes, o)
			}
			return o
		} else if !strict {
			return r
		}
		return (-1)
	}, s)
}

// Decipher a message from ciphertext to plaintext.
func (c VigenereFamilyCipher) Decipher(s string, strict bool) string {
	keyRunes := []rune(c.countersign)
	var transcodedCharCount = 0
	return strings.Map(func(r rune) rune {
		k := keyRunes[transcodedCharCount%len(keyRunes)]
		cipher := c.ciphers[k]
		if o := cipher.decipherRune(r); o != (-1) {
			transcodedCharCount++
			if c.Textautoclave {
				keyRunes = append(keyRunes, o)
			} else if c.Keyautoclave {
				keyRunes = append(keyRunes, r)
			}
			return o
		} else if !strict {
			return r
		}
		return (-1)
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

// MakeVigenereCipher creates a new Vigenere cipher.
func MakeVigenereCipher(countersign, alphabet string) VigenereFamilyCipher {
	return MakeVigenereFamilyCipher(countersign, alphabet, alphabet, alphabet)
}

// MakeVigenereTextAutoclaveCipher creates a new Vigenere (text autoclave) cipher.
func MakeVigenereTextAutoclaveCipher(countersign, alphabet string) VigenereFamilyCipher {
	c := MakeVigenereCipher(countersign, alphabet)
	c.Textautoclave = true
	return c
}

// MakeVigenereKeyAutoclaveCipher creates a new Vigenere (key autoclave) cipher.
func MakeVigenereKeyAutoclaveCipher(countersign, alphabet string) VigenereFamilyCipher {
	c := MakeVigenereCipher(countersign, alphabet)
	c.Keyautoclave = true
	return c
}

// MakeBeaufortCipher creates a new Beaufort cipher.
func MakeBeaufortCipher(countersign, alphabet string) VigenereFamilyCipher {
	revAlphabet := reverseString(alphabet)
	return MakeVigenereFamilyCipher(countersign, alphabet, revAlphabet, revAlphabet)
}

// MakeGronsfeldCipher creates a new Gronsfeld cipher.
func MakeGronsfeldCipher(countersign, alphabet string) VigenereFamilyCipher {
	return MakeVigenereFamilyCipher(countersign, alphabet, alphabet, "0123456789")
}

// MakeVariantBeaufortCipher creates a new Vigenere cipher.
func MakeVariantBeaufortCipher(countersign, alphabet string) VigenereFamilyCipher {
	revAlphabet := reverseString(alphabet)
	return MakeVigenereFamilyCipher(countersign, revAlphabet, revAlphabet, alphabet)
}

// MakeTrithemiusCipher creates a new Trithemius cipher.
// MakeTrithemiusCipher considers this simply the Vigenere cipher with the countersign equal to the alphabet.
func MakeTrithemiusCipher(alphabet string) VigenereFamilyCipher {
	return MakeVigenereCipher(alphabet, alphabet)
}

// MakeDellaPortaCipher creates a new DellaPorta cipher.
func MakeDellaPortaCipher(countersign, alphabet string) VigenereFamilyCipher {
	return MakeDellaPortaReciprocalTable(countersign, alphabet, alphabet, alphabet)
}
