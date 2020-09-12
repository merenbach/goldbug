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
	"text/tabwriter"

	"github.com/merenbach/goldbug/internal/masc"
)

// ReciprocalTable holds a reciprocal table.
type ReciprocalTable struct {
	Strict bool

	KeyAlphabet string
	PtAlphabet  string
	CtAlphabets []string
}

func makedicts(columnHeaders string, rowHeaders string, rows []string) (map[rune]*masc.Tableau, error) {
	pt2ct := make(map[rune]*masc.Tableau)

	keyRunes := []rune(rowHeaders)
	if len(keyRunes) != len(rowHeaders) {
		return nil, errors.New("Row headers must have same rune length as rows slice")
	}

	for i, r := range keyRunes {
		pt2ct[r] = &masc.Tableau{
			PtAlphabet: columnHeaders,
			CtAlphabet: rows[i],
			Strict:     true,
		}
	}

	return pt2ct, nil
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
	var b strings.Builder

	w := tabwriter.NewWriter(&b, 4, 1, 3, ' ', 0)

	formatForPrinting := func(s string) string {
		spl := strings.Split(s, "")
		return strings.Join(spl, " ")
	}

	fmt.Fprintf(w, "\t%s\n", formatForPrinting(tr.PtAlphabet))
	for i, r := range []rune(tr.KeyAlphabet) {
		fmt.Fprintf(w, "\n%c\t%s", r, formatForPrinting(tr.CtAlphabets[i]))
	}

	w.Flush()
	return b.String(), nil
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
	pt2ct, err := makedicts(tr.PtAlphabet, tr.KeyAlphabet, tr.CtAlphabets)
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

		if o := m.EncipherRune(r); o != (-1) {
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
	ct2pt, err := makedicts(tr.PtAlphabet, tr.KeyAlphabet, tr.CtAlphabets)
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

		if o := m.DecipherRune(r); o != (-1) {
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
