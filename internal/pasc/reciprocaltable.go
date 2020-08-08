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

import (
	"errors"
	"fmt"
	"strings"

	"github.com/merenbach/goldbug/internal/translation"
)

// ReciprocalTable holds a reciprocal table.
type ReciprocalTable struct {
	Strict bool

	KeyAlphabet string
	PtAlphabet  string
	CtAlphabets []string
}

func makedicts(columnHeaders string, rowHeaders string, rows []string) (map[rune]map[rune]rune, map[rune]map[rune]rune, error) {
	pt2ct := make(map[rune]map[rune]rune)
	ct2pt := make(map[rune]map[rune]rune)

	keyRunes := []rune(rowHeaders)
	if len(keyRunes) != len(rowHeaders) {
		return nil, nil, errors.New("Row headers must have same rune length as rows slice")
	}

	for i, r := range keyRunes {
		ptTable := translation.Table{
			Src: columnHeaders,
			Dst: rows[i],
		}
		ptCipher, err := ptTable.Map()
		if err != nil {
			return nil, nil, err
		}

		ctTable := translation.Table{
			Src: rows[i],
			Dst: columnHeaders,
		}
		ctCipher, err := ctTable.Map()
		if err != nil {
			return nil, nil, err
		}

		pt2ct[r], ct2pt[r] = ptCipher, ctCipher
	}

	return pt2ct, ct2pt, nil
}

func (tr *ReciprocalTable) String() string {
	return fmt.Sprintf("%+v", map[string]interface{}{
		"k":  tr.KeyAlphabet,
		"pt": tr.PtAlphabet,
		"ct": tr.CtAlphabets,
	})
}

// Printable representation of this tabula recta.
func (tr *ReciprocalTable) Printable() (string, error) {
	var out strings.Builder
	var err error
	formatForPrinting := func(s string) string {
		spl := strings.Split(s, "")
		return strings.Join(spl, " ")
	}
	if _, err = out.WriteString("    " + formatForPrinting(tr.PtAlphabet) + "\n  +"); err != nil {
		return "", err
	}
	for range tr.PtAlphabet {
		if _, err = out.WriteRune('-'); err != nil {
			return "", err
		}
		if _, err = out.WriteRune('-'); err != nil {
			return "", err
		}
	}
	for i, r := range tr.KeyAlphabet {
		ctAlpha := fmt.Sprintf("\n%c | %s", r, formatForPrinting(tr.CtAlphabets[i]))
		if _, err = out.WriteString(ctAlpha); err != nil {
			return "", err
		}
	}
	return out.String(), nil
}

// // Encipher a plaintext rune with a given key alphabet rune.
// // Encipher will return (-1, false) if the key rune is invalid.
// // Encipher will return (-1, true) if the key rune is valid but the message rune is not.
// // Encipher will otherwise return the transcoded rune as the first argument and true as the second.
// func (tr *TabulaRecta) encipher(r rune, k rune) (rune, bool) {
// 	c, ok := tr.pt2ct[k]
// 	if !ok {
// 		return (-1), false
// 	}

// 	if o, ok := c[r]; ok {
// 		return o, true
// 	}
// 	return (-1), true
// }

// // Decipher a ciphertext rune with a given key alphabet rune.
// // Decipher will return (-1, false) if the key rune is invalid.
// // Decipher will return (-1, true) if the key rune is valid but the message rune is not.
// // Decipher will otherwise return the transcoded rune as the first argument and true as the second.
// func (tr *TabulaRecta) decipher(r rune, k rune) (rune, bool) {
// 	c, ok := tr.ct2pt[k]
// 	if !ok {
// 		return (-1), false
// 	}

// 	if o, ok := c[r]; ok {
// 		return o, true
// 	}
// 	return (-1), true
// }

// Encipher a string.
// Encipher will invoke the onSuccess function with before and after runes.
func (tr *ReciprocalTable) Encipher(s string, k string, onSuccess func(rune, rune, *[]rune)) (string, error) {
	pt2ct, _, err := makedicts(tr.PtAlphabet, tr.KeyAlphabet, tr.CtAlphabets)
	if err != nil {
		return "", err
	}

	keyRunes := []rune(k)
	var transcodedCharCount = 0
	return strings.Map(func(r rune) rune {
		k := keyRunes[transcodedCharCount%len(keyRunes)]
		m, ok := pt2ct[k]
		if !ok {
			// Rune `k` does not exist in keyAlphabet
			// TODO: avoid advancing on invalid key char
			// TODO: avoid infinite loop upon _no_ valid key chars
			return (-1)
		}

		if o, ok := m[r]; ok {
			// Transcoding successful
			transcodedCharCount++
			if onSuccess != nil {
				onSuccess(r, o, &keyRunes)
			}
			return o
		} else if !tr.Strict {
			// Rune `r` does not exist in ptAlphabet but we are not being struct
			return r
		}
		// Rune `r` does not exist in ptAlphabet
		return (-1)
	}, s), nil
}

// Decipher a string.
// Decipher will invoke the onSuccess function with before and after runes.
func (tr *ReciprocalTable) Decipher(s string, k string, onSuccess func(rune, rune, *[]rune)) (string, error) {
	_, ct2pt, err := makedicts(tr.PtAlphabet, tr.KeyAlphabet, tr.CtAlphabets)
	if err != nil {
		return "", err
	}

	keyRunes := []rune(k)
	var transcodedCharCount = 0
	return strings.Map(func(r rune) rune {
		k := keyRunes[transcodedCharCount%len(keyRunes)]
		m, ok := ct2pt[k]
		if !ok {
			// Rune `k` does not exist in keyAlphabet
			// TODO: avoid advancing on invalid key char
			// TODO: avoid infinite loop upon _no_ valid key chars
			return (-1)
		}

		if o, ok := m[r]; ok {
			// Transcoding successful
			transcodedCharCount++
			if onSuccess != nil {
				onSuccess(r, o, &keyRunes)
			}
			return o
		} else if !tr.Strict {
			// Rune `r` does not exist in ctAlphabet but we are not being strict
			return r
		}
		// Rune `r` does not exist in ctAlphabet
		return (-1)
	}, s), nil
}
