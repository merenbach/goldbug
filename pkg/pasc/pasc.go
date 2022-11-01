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
	"unicode"

	"github.com/merenbach/goldbug/internal/sliceutil"
	"github.com/merenbach/goldbug/pkg/masc"
)

type autokeyFunc func(rune, rune, rune) rune

// A TabulaRecta holds a tabula recta.
type tabulaRecta struct {
	// this is used here only for key lookups, hence no corresponding "strict" option at this time
	caseless bool

	ptAlphabet string
	// CtAlphabet  string
	keyAlphabet string

	key string

	autokeyer autokeyFunc

	tableau map[rune]*masc.SimpleCipher
}

func newTabulaRecta(ptAlphabet string, keyAlphabet string, key string, ciphers []*masc.SimpleCipher, autokeyer autokeyFunc, caseless bool) (*tabulaRecta, error) {
	keyAlphabetRunes := []rune(keyAlphabet)
	if len(keyAlphabetRunes) != len(ciphers) {
		return nil, errors.New("row headers must have same rune length as rows slice")
	}

	tableau, err := sliceutil.Zipmap(keyAlphabetRunes, ciphers)
	if err != nil {
		return nil, fmt.Errorf("could not zip alphabet with ciphers: %w", err)
	}

	return &tabulaRecta{
		ptAlphabet:  ptAlphabet,
		keyAlphabet: keyAlphabet,
		key:         key,
		autokeyer:   autokeyer,
		caseless:    caseless,
		tableau:     tableau,
	}, nil
}

// Printable representation of this tabula recta.
func (tr *tabulaRecta) Printable() (string, error) {
	ptAlphabet := tr.ptAlphabet

	keyAlphabet := tr.keyAlphabet
	if keyAlphabet == "" {
		keyAlphabet = ptAlphabet
	}

	rt := tr.tableau

	var b strings.Builder

	w := tabwriter.NewWriter(&b, 4, 1, 3, ' ', 0)

	formatForPrinting := func(s string) string {
		spl := strings.Split(s, "")
		return strings.Join(spl, " ")
	}

	fmt.Fprintf(w, "\t%s\n", formatForPrinting(ptAlphabet))
	for _, r := range keyAlphabet {
		if tableau, ok := rt[r]; ok {
			out, err := tableau.Encipher(ptAlphabet)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(w, "\n%c\t%s", r, formatForPrinting(out))
		}
	}

	w.Flush()
	return b.String(), nil
}

// // Encipher a plaintext rune with a given key alphabet rune.
// // Encipher will return (-1, false) if the key rune is invalid.
// // Encipher will return (-1, true) if the key rune is valid but the message rune is not.
// // Encipher will otherwise return the transcoded rune as the first argument and true as the second.
// func (tr *tabulaRecta) encipher(r rune, k rune) (rune, bool) {
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
// func (tr *tabulaRecta) decipher(r rune, k rune) (rune, bool) {
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
func (tr *tabulaRecta) Encipher(s string) (string, error) {
	if len(tr.key) == 0 {
		return s, nil
	}

	tableau := tr.tableau

	keyRunes := []rune(tr.key)

	var transcodedCharCount = 0
	return strings.Map(func(r rune) rune {
		k := keyRunes[transcodedCharCount%len(keyRunes)]
		m, ok := tableau[k]
		if !ok && tr.caseless {
			m, ok = tableau[unicode.ToUpper(k)]
		}
		if !ok && tr.caseless {
			m, ok = tableau[unicode.ToLower(k)]
		}
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
			if tr.autokeyer != nil {
				keyRunes = append(keyRunes, tr.autokeyer(k, r, o))
			}
		}
		return o
	}, s), nil
}

// Decipher a string.
// Decipher will invoke the onSuccess function with before and after runes.
func (tr *tabulaRecta) Decipher(s string) (string, error) {
	if len(tr.key) == 0 {
		return s, nil
	}

	tableau := tr.tableau

	keyRunes := []rune(tr.key)

	var transcodedCharCount = 0
	return strings.Map(func(r rune) rune {
		k := keyRunes[transcodedCharCount%len(keyRunes)]
		m, ok := tableau[k]
		if !ok && tr.caseless {
			m, ok = tableau[unicode.ToUpper(k)]
		}
		if !ok && tr.caseless {
			m, ok = tableau[unicode.ToLower(k)]
		}
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
			if tr.autokeyer != nil {
				keyRunes = append(keyRunes, tr.autokeyer(k, o, r))
			}
		}
		return o
	}, s), nil
}
