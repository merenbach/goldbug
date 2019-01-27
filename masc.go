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

func makeSimpleSubstitutionCipher(ptAlphabet string, ctAlphabet string) SimpleSubstitutionCipher {
	return SimpleSubstitutionCipher{
		ptAlphabet: ptAlphabet,
		ctAlphabet: ctAlphabet,
		pt2ct:      makeRuneMap(ptAlphabet, ctAlphabet),
		ct2pt:      makeRuneMap(ctAlphabet, ptAlphabet),
	}
}

// MakeRuneMap creates a one-way, monoalphabetic substitution cipher table.
func makeRuneMap(src, dst string) map[rune]rune {
	out := make(map[rune]rune)

	dstRunes := []rune(dst)
	for i, r := range []rune(src) {
		out[r] = dstRunes[i]
	}

	return out
}

func (c SimpleSubstitutionCipher) String() string {
	return fmt.Sprintf("PT: %s\nCT: %s", c.ptAlphabet, c.ctAlphabet)
}

// Encipher a message from plaintext to ciphertext.
func (c SimpleSubstitutionCipher) Encipher(s string, strict bool) string {
	return strings.Map(func(r rune) rune {
		if o, ok := c.pt2ct[r]; ok {
			return o
		} else if !strict {
			return r
		}
		return (-1)
	}, s)
}

// Decipher a message from ciphertext to plaintext.
func (c SimpleSubstitutionCipher) Decipher(s string, strict bool) string {
	return strings.Map(func(r rune) rune {
		if o, ok := c.ct2pt[r]; ok {
			return o
		} else if !strict {
			return r
		}
		return (-1)
	}, s)
}

// EncipherRune transforms a rune from plaintext to ciphertext, returning (-1) if transformation fails.
func (c SimpleSubstitutionCipher) encipherRune(r rune) rune {
	if o, ok := c.pt2ct[r]; ok {
		return o
	}
	return (-1)
}

// DecipherRune transforms a rune from ciphertext to plaintext, returning (-1) if transformation fails.
func (c SimpleSubstitutionCipher) decipherRune(r rune) rune {
	if o, ok := c.ct2pt[r]; ok {
		return o
	}
	return (-1)
}

// NewKeywordCipher creates a new keyword cipher.
func MakeKeywordCipher(alphabet, keyword string) SimpleSubstitutionCipher {
	ctAlphabet := deduplicate(keyword + alphabet)
	return makeSimpleSubstitutionCipher(alphabet, ctAlphabet)
}

// NewAffineCipher creates a new affine cipher.
func MakeAffineCipher(ptAlphabet string, a, b int) SimpleSubstitutionCipher {
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

	return makeSimpleSubstitutionCipher(ptAlphabet, ctAlphabet)
}

// NewAtbashCipher creates a new Atbash cipher.
func MakeAtbashCipher(ptAlphabet string) SimpleSubstitutionCipher {
	return MakeAffineCipher(ptAlphabet, -1, -1)
}

// NewCaesarCipher creates a new Caesar cipher.
func MakeCaesarCipher(ptAlphabet string, b int) SimpleSubstitutionCipher {
	return MakeAffineCipher(ptAlphabet, 1, b)
}

// NewDecimationCipher creates a new decimation cipher.
func MakeDecimationCipher(ptAlphabet string, a int) SimpleSubstitutionCipher {
	return MakeAffineCipher(ptAlphabet, a, 0)
}

// NewRot13Cipher creates a new Rot13 (Caesar shift of 13) cipher.
func MakeRot13Cipher(ptAlphabet string) SimpleSubstitutionCipher {
	return MakeCaesarCipher(ptAlphabet, 13)
}
