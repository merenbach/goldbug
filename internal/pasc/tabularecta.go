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

	"github.com/merenbach/goldbug/pkg/masc"
)

// TabulaRecta holds a tabula recta.
type TabulaRecta struct {
	caseless bool

	ptAlphabet string
	// CtAlphabet  string
	keyAlphabet string

	key string

	dictFunc  func(string, int) (*masc.SimpleCipher, error)
	autokeyer func(rune, rune, *[]rune)

	tableau map[rune]*masc.SimpleCipher
}

// adapted from: https://www.sohamkamani.com/golang/options-pattern/

type TabulaRectaOption func(*TabulaRecta)

// func WithStrict(b bool) TabulaRectaOption {
// 	return func(c *TabulaRecta) {
// 		c.strict = b
// 	}
// }

func WithCaseless(b bool) TabulaRectaOption {
	return func(c *TabulaRecta) {
		c.caseless = b
	}
}

func WithPtAlphabet(s string) TabulaRectaOption {
	return func(c *TabulaRecta) {
		c.ptAlphabet = s
	}
}

// func WithCtAlphabet(s string) TableauOption {
// 	return func(c *TabulaRecta) {
// 		if s != "" {
// 			c.ctAlphabet = s
// 		}
// 	}
// }

func WithKeyAlphabet(s string) TabulaRectaOption {
	return func(c *TabulaRecta) {
		c.keyAlphabet = s
	}
}

func WithKey(s string) TabulaRectaOption {
	return func(c *TabulaRecta) {
		c.key = s
	}
}

func WithAutokeyer(f func(rune, rune, *[]rune)) TabulaRectaOption {
	return func(c *TabulaRecta) {
		c.autokeyer = f
	}
}

func WithDictFunc(f func(s string, i int) (*masc.SimpleCipher, error)) TabulaRectaOption {
	return func(c *TabulaRecta) {
		c.dictFunc = f
	}
}

func NewTabulaRecta(opts ...TabulaRectaOption) (*TabulaRecta, error) {
	t := &TabulaRecta{
		ptAlphabet:  Alphabet,
		keyAlphabet: Alphabet,
		// ctAlphabet: Alphabet,
	}
	for _, opt := range opts {
		opt(t)
	}

	tableau, err := t.maketableau()
	if err != nil {
		return nil, fmt.Errorf("could not create tableau: %w", err)
	}

	t.tableau = tableau

	return t, nil
}

func (tr *TabulaRecta) maketableau() (map[rune]*masc.SimpleCipher, error) {
	f := tr.dictFunc

	ptAlphabet, keyAlphabet := tr.ptAlphabet, tr.keyAlphabet

	if keyAlphabet == "" {
		keyAlphabet = ptAlphabet
	}

	m := make(map[rune]*masc.SimpleCipher)

	keyRunes := []rune(keyAlphabet)

	if len(keyRunes) != len(keyAlphabet) {
		return nil, errors.New("row headers must have same rune length as rows slice")
	}

	for i, r := range keyRunes {
		t, err := f(ptAlphabet, i)
		if err != nil {
			return nil, err
		}
		m[r] = t
	}

	return m, nil
}

// Printable representation of this tabula recta.
func (tr *TabulaRecta) Printable() (string, error) {
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
func (tr *TabulaRecta) Encipher(s string) (string, error) {
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
				tr.autokeyer(r, o, &keyRunes)
			}
		}
		return o
	}, s), nil
}

// Decipher a string.
// Decipher will invoke the onSuccess function with before and after runes.
func (tr *TabulaRecta) Decipher(s string) (string, error) {
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
				tr.autokeyer(o, r, &keyRunes)
			}
		}
		return o
	}, s), nil
}
