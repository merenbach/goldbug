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

package masc

// Monoalphabetic substitution ciphers

import (
	"github.com/merenbach/gold-bug/internal/stringutil"
)

// DefaultAlphabet is the default character set for monoalphabetic substitution ciphers
const defaultAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// NewKeywordCipher creates a new keyword cipher.
// NewKeywordCipher currently doesn't return any errors, but returns the error argument for shape consistency.
func NewKeywordCipher(ptAlphabet string, keyword string) (*MonoalphabeticSubstitutionCipher, error) {
	if ptAlphabet == "" {
		ptAlphabet = defaultAlphabet
	}

	ctAlphabet := stringutil.Deduplicate(keyword + ptAlphabet)
	return New(ptAlphabet, ctAlphabet)
}

// NewAffineCipher creates a new affine cipher.
func NewAffineCipher(ptAlphabet string, a int, b int) (*MonoalphabeticSubstitutionCipher, error) {
	if ptAlphabet == "" {
		ptAlphabet = defaultAlphabet
	}

	ctAlphabet, err := stringutil.Affine(ptAlphabet, a, b)
	if err != nil {
		return nil, err
	}

	return New(ptAlphabet, ctAlphabet)
}

// NewAtbashCipher creates a new Atbash cipher.
func NewAtbashCipher(ptAlphabet string) (*MonoalphabeticSubstitutionCipher, error) {
	return NewAffineCipher(ptAlphabet, -1, -1)
}

// NewCaesarCipher creates a new Caesar cipher.
func NewCaesarCipher(ptAlphabet string, b int) (*MonoalphabeticSubstitutionCipher, error) {
	return NewAffineCipher(ptAlphabet, 1, b)
}

// NewDecimationCipher creates a new decimation cipher.
func NewDecimationCipher(ptAlphabet string, a int) (*MonoalphabeticSubstitutionCipher, error) {
	return NewAffineCipher(ptAlphabet, a, 0)
}

// NewRot13Cipher creates a new Rot13 (Caesar shift of 13) cipher.
func NewRot13Cipher(ptAlphabet string) (*MonoalphabeticSubstitutionCipher, error) {
	return NewCaesarCipher(ptAlphabet, 13)
}
