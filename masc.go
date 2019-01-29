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

// Monoalphabetic substitution ciphers

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// A SimpleSubstitutionCipher represents a simple monoalphabetic substitution cipher.
type SimpleSubstitutionCipher struct {
	ptAlphabet string
	ctAlphabet string
	pt2ct      map[rune]rune
	ct2pt      map[rune]rune
}

// NewSimpleSubstitutionCipher returns a basic monoalphabetic substitution cipher with the given alphabet translations.
func NewSimpleSubstitutionCipher(ptAlphabet string, ctAlphabet string) *SimpleSubstitutionCipher {
	return &SimpleSubstitutionCipher{
		ptAlphabet: ptAlphabet,
		ctAlphabet: ctAlphabet,
		pt2ct:      makeRuneMap(ptAlphabet, ctAlphabet),
		ct2pt:      makeRuneMap(ctAlphabet, ptAlphabet),
	}
}

// MakeRuneMap creates a one-way, monoalphabetic substitution cipher table.
// TODO: pass in as runes? error out on duplicates?
func makeRuneMap(src, dst string) map[rune]rune {
	out := make(map[rune]rune)

	dstRunes := []rune(dst)
	for i, r := range []rune(src) {
		out[r] = dstRunes[i]
	}

	return out
}

func (sc *SimpleSubstitutionCipher) String() string {
	return fmt.Sprintf("PT: %s\nCT: %s", sc.ptAlphabet, sc.ctAlphabet)
}

// Encipher a message from plaintext to ciphertext.
func (sc *SimpleSubstitutionCipher) Encipher(s string, strict bool) string {
	return strings.Map(func(r rune) rune {
		o, _ := sc.EncipherRune(r, strict)
		return o
	}, s)
}

// Decipher a message from ciphertext to plaintext.
func (sc *SimpleSubstitutionCipher) Decipher(s string, strict bool) string {
	return strings.Map(func(r rune) rune {
		o, _ := sc.DecipherRune(r, strict)
		return o
	}, s)
}

// EncipherRune transforms a rune from plaintext to ciphertext, returning (-1, false) if transformation fails.
func (sc *SimpleSubstitutionCipher) encipherRune(r rune) (rune, bool) {
	if o, ok := sc.pt2ct[r]; ok {
		return o, true
	}
	return (-1), false
}

// DecipherRune transforms a rune from ciphertext to plaintext, returning (-1, false) if transformation fails.
func (sc *SimpleSubstitutionCipher) decipherRune(r rune) (rune, bool) {
	if o, ok := sc.ct2pt[r]; ok {
		return o, true
	}
	return (-1), false
}

// EncipherRune transforms a rune from plaintext to ciphertext, returning (-1, false) if transformation fails and strict mode is false.
// EncipherRune returns a second parameter, a boolean, to indicate if any transformation was possible.
func (sc *SimpleSubstitutionCipher) EncipherRune(r rune, strict bool) (rune, bool) {
	if o, ok := sc.encipherRune(r); ok || strict {
		return o, ok
	}
	return r, false
}

// DecipherRune transforms a rune from ciphertext to plaintext, returning (-1, false) if transformation fails and strict mode is false.
// DecipherRune returns a second parameter, a boolean, to indicate if any transformation was possible.
func (sc *SimpleSubstitutionCipher) DecipherRune(r rune, strict bool) (rune, bool) {
	if o, ok := sc.decipherRune(r); ok || strict {
		return o, ok
	}
	return r, false
}

// NewKeywordCipher creates a new keyword cipher.
func NewKeywordCipher(alphabet, keyword string) *SimpleSubstitutionCipher {
	ctAlphabet := deduplicate(keyword + alphabet)
	return NewSimpleSubstitutionCipher(alphabet, ctAlphabet)
}

// NewAffineCipher creates a new affine cipher.
func NewAffineCipher(ptAlphabet string, a, b int) *SimpleSubstitutionCipher {
	m := utf8.RuneCountInString(ptAlphabet)

	// TODO: consider using Hull-Dobell satisfaction to determine if `a` is valid (must be coprime with `m`)
	for a < 0 {
		a += m
	}
	for b < 0 {
		b += m
	}

	lcg := LCG{
		Modulus:    uint(m),
		Multiplier: 1,
		Increment:  uint(a),
		Seed:       uint(b),
	}
	ctAlphabet := backpermute(ptAlphabet, lcg.Iterator())

	return NewSimpleSubstitutionCipher(ptAlphabet, ctAlphabet)
}

// NewAtbashCipher creates a new Atbash cipher.
func NewAtbashCipher(ptAlphabet string) *SimpleSubstitutionCipher {
	return NewAffineCipher(ptAlphabet, -1, -1)
}

// NewCaesarCipher creates a new Caesar cipher.
func NewCaesarCipher(ptAlphabet string, b int) *SimpleSubstitutionCipher {
	return NewAffineCipher(ptAlphabet, 1, b)
}

// NewDecimationCipher creates a new decimation cipher.
func NewDecimationCipher(ptAlphabet string, a int) *SimpleSubstitutionCipher {
	return NewAffineCipher(ptAlphabet, a, 0)
}

// NewRot13Cipher creates a new Rot13 (Caesar shift of 13) cipher.
func NewRot13Cipher(ptAlphabet string) *SimpleSubstitutionCipher {
	return NewCaesarCipher(ptAlphabet, 13)
}
