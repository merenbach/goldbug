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

import (
	"fmt"

	"github.com/merenbach/gold-bug/internal/stringutil"
)

// TODO: worth translating on the fly and memoizing to allow creation with bare struct?
//       but how would we allow struct updates to propagate to the memoized data?

// A MonoalphabeticSubstitutionCipher represents a simple monoalphabetic substitution cipher.
type MonoalphabeticSubstitutionCipher struct {
	ptAlphabet string
	ctAlphabet string

	pt2ct map[rune]rune // map[pt]=>ct
	ct2pt map[rune]rune // map[ct]=>pt

	// Strict mode removes characters that cannot be translated
	Strict bool
}

// New monoalphabetic substitution cipher with the given plaintext and ciphertext alphabets.
func New(pt string, ct string) (*MonoalphabeticSubstitutionCipher, error) {
	pt2ct, err := stringutil.MakeTranslationTable(pt, ct, "")
	if err != nil {
		return nil, err
	}

	ct2pt, err := stringutil.MakeTranslationTable(ct, pt, "")
	if err != nil {
		return nil, err
	}

	return &MonoalphabeticSubstitutionCipher{
		ptAlphabet: pt,
		ctAlphabet: ct,
		pt2ct:      pt2ct,
		ct2pt:      ct2pt,
	}, nil
}

func (sc *MonoalphabeticSubstitutionCipher) String() string {
	return fmt.Sprintf("%+v", map[string]interface{}{
		"pt": sc.ptAlphabet,
		"ct": sc.ctAlphabet,
	})
}

// Printable representation of this tableau.
func (sc *MonoalphabeticSubstitutionCipher) Printable() string {
	return fmt.Sprintf("PT: %s\nCT: %s", sc.ptAlphabet, sc.ctAlphabet)
}

// EncipherString transforms a message from plaintext to ciphertext.
func (sc *MonoalphabeticSubstitutionCipher) EncipherString(s string) (string, error) {
	return stringutil.Translate(s, sc.pt2ct, sc.Strict), nil
}

// DecipherString transforms a message from ciphertext to plaintext.
func (sc *MonoalphabeticSubstitutionCipher) DecipherString(s string) (string, error) {
	return stringutil.Translate(s, sc.ct2pt, sc.Strict), nil
}
