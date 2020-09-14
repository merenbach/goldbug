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
	"strings"

	"github.com/merenbach/goldbug/internal/masc"
)

type ReciprocalTable map[rune]*masc.Tableau

// func makedicts(columnHeaders string, rowHeaders string, rows []string, strict bool) (map[rune]*masc.Tableau, error) {
// 	m := make(map[rune]*masc.Tableau)

// 	keyRunes := []rune(rowHeaders)
// 	if len(keyRunes) != len(rowHeaders) {
// 		return nil, errors.New("Row headers must have same rune length as rows slice")
// 	}

// 	for i, r := range keyRunes {
// 		t, err := masc.New(columnHeaders, func(string) (string, error) {
// 			return rows[i], nil
// 		})
// 		if err != nil {
// 			return nil, err
// 		}
// 		t.Strict = strict
// 		m[r] = t
// 	}

// 	return m, nil
// }

// func (tr *ReciprocalTable) String() string {
// 	return fmt.Sprintf("%+v", map[string]interface{}{
// 		"k":  tr.KeyAlphabet,
// 		"pt": tr.PtAlphabet,
// 		"ct": tr.CtAlphabets,
// 	})
// }

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
func (tr ReciprocalTable) Encipher(s string, k string, onSuccess func(rune, rune, *[]rune)) (string, error) {
	keyRunes := []rune(k)
	var transcodedCharCount = 0
	return strings.Map(func(r rune) rune {
		k := keyRunes[transcodedCharCount%len(keyRunes)]
		m, ok := tr[k]
		if !ok {
			// Rune `k` does not exist in keyAlphabet
			// TODO: avoid advancing on invalid key char
			// TODO: avoid infinite loop upon _no_ valid key chars
			return (-1)
		}

		o, ok := m.EncipherRune(r)
		if ok {
			// Transcoding successful
			transcodedCharCount++
			if onSuccess != nil {
				onSuccess(r, o, &keyRunes)
			}
		}
		return o
	}, s), nil
}

// Decipher a string.
// Decipher will invoke the onSuccess function with before and after runes.
func (tr ReciprocalTable) Decipher(s string, k string, onSuccess func(rune, rune, *[]rune)) (string, error) {
	keyRunes := []rune(k)
	var transcodedCharCount = 0
	return strings.Map(func(r rune) rune {
		k := keyRunes[transcodedCharCount%len(keyRunes)]
		m, ok := tr[k]
		if !ok {
			// Rune `k` does not exist in keyAlphabet
			// TODO: avoid advancing on invalid key char
			// TODO: avoid infinite loop upon _no_ valid key chars
			return (-1)
		}

		o, ok := m.DecipherRune(r)
		if ok {
			// Transcoding successful
			transcodedCharCount++
			if onSuccess != nil {
				onSuccess(r, o, &keyRunes)
			}
		}
		return o
	}, s), nil
}
