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

package pasc

// Polyalphabetic substitution ciphers

import (
	"errors"
	"strings"
	"unicode/utf8"

	"github.com/merenbach/goldbug/internal/stringutil"
	"github.com/merenbach/goldbug/internal/tabularecta"
)

// Alphabet to use by default for substitution ciphers
const Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// A VigenereFamilyCipher represents a cipher in the Vigenere family.
type VigenereFamilyCipher struct {
	*tabularecta.TabulaRecta
	countersign string

	// Strict mode removes characters that cannot be translated
	Strict bool

	// TextAutoclave mode appends the plaintext to the key during the encoding process
	TextAutoclave bool

	// KeyAutoclave mode appends the ciphertext to the key during the encoding process
	KeyAutoclave bool
}

// NewVigenereFamilyCipher creates a new tabula recta suitable for use with the Vigenere family of ciphers.
// NOTE: we roll the countersign into the tabula recta so it has all the data it needs
// to decode/encode a string reusably, for parallelism with the monoalphabetic ciphers.
func NewVigenereFamilyCipher(countersign string, ptAlphabet string, ctAlphabet string, keyAlphabet string) (*VigenereFamilyCipher, error) {
	keyRunes := []rune(keyAlphabet)
	ctAlphabets := make([]string, len(keyRunes))

	// Cast to increase index without gaps
	for i := range keyRunes {
		// ctAlphabetRotated, err := stringutil.Affine(ctAlphabet, 1, i)
		// if err != nil {
		// 	return nil
		// }
		ctAlphabets[i] = stringutil.WrapString(ctAlphabet, i)
	}

	tr, err := tabularecta.New(ptAlphabet, keyAlphabet, ctAlphabets)
	if err != nil {
		return nil, err
	}

	return &VigenereFamilyCipher{
		TabulaRecta: tr,
		countersign: countersign,
	}, nil
}

// NewDellaPortaReciprocalTable creates a new tabula recta suitable for use with the Della Porta cipher.
func NewDellaPortaReciprocalTable(countersign string, ptAlphabet string, ctAlphabet string, keyAlphabet string) (*VigenereFamilyCipher, error) {
	keyRunes := []rune(keyAlphabet)
	ctAlphabets := make([]string, len(keyRunes))

	if utf8.RuneCountInString(ctAlphabet)%2 != 0 {
		return nil, errors.New("Della Porta cipher alphabets must have even length")
	}
	ctAlphabetWrapped := stringutil.WrapString(ctAlphabet, utf8.RuneCountInString(ctAlphabet)/2)

	for i := range keyRunes {
		ctAlphabets[i] = owrapString(ctAlphabetWrapped, i/2)
	}

	tr, err := tabularecta.New(ptAlphabet, keyAlphabet, ctAlphabets)
	if err != nil {
		return nil, err
	}

	return &VigenereFamilyCipher{
		TabulaRecta: tr,
		countersign: countersign,
	}, nil
}

// EncipherString transforms a string from plaintext to ciphertext.
func (sc *VigenereFamilyCipher) EncipherString(s string) string {
	keyRunes := []rune(sc.countersign)
	var transcodedCharCount = 0
	return strings.Map(func(r rune) rune {
		k := keyRunes[transcodedCharCount%len(keyRunes)]
		o, ok := sc.Encipher(r, k)
		if !ok {
			// TODO: avoid advancing on invalid key char
			// TODO: avoid infinite loop upon _no_ valid key chars
			return (-1)
		}
		if o != -1 {
			transcodedCharCount++
			if sc.TextAutoclave {
				keyRunes = append(keyRunes, r)
			} else if sc.KeyAutoclave {
				keyRunes = append(keyRunes, o)
			}
			return o
		} else if !sc.Strict {
			return r
		}
		return (-1)
	}, s)
}

// DecipherString transforms a string from ciphertext to plaintext.
func (sc *VigenereFamilyCipher) DecipherString(s string) string {
	keyRunes := []rune(sc.countersign)
	var transcodedCharCount = 0
	return strings.Map(func(r rune) rune {
		k := keyRunes[transcodedCharCount%len(keyRunes)]
		o, ok := sc.Decipher(r, k)
		if !ok {
			// TODO: avoid advancing on invalid key char
			// TODO: avoid infinite loop upon _no_ valid key chars
			return (-1)
		}
		if o != -1 {
			transcodedCharCount++
			if sc.TextAutoclave {
				keyRunes = append(keyRunes, o)
			} else if sc.KeyAutoclave {
				keyRunes = append(keyRunes, r)
			}
			return o
		} else if !sc.Strict {
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
	return stringutil.WrapString(string(u), i) + stringutil.WrapString(string(v), len(v)-i)
}
