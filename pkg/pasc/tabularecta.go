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
	"fmt"
	"strings"
	"text/tabwriter"
	"unicode"
	"unicode/utf8"

	"github.com/merenbach/goldbug/internal/sliceutil"
	"github.com/merenbach/goldbug/pkg/masc"
)

type autokeyFunc func(rune, rune, rune) rune

// No autokey
var defaultAutokeyFunc autokeyFunc = func(k rune, _ rune, _ rune) rune {
	return k
}

// Key autokey
var keyAutokeyFunc autokeyFunc = func(_ rune, a rune, _ rune) rune {
	return a
}

// Text autokey
var textAutokeyFunc autokeyFunc = func(_ rune, _ rune, b rune) rune {
	return b
}

// An AutokeyOption determines the autokey setting for the cipher.
type autokeyOption uint8

const (
	// NoAutokey signifies not to autokey the cipher.
	NoAutokey autokeyOption = iota

	// TextAutokey denotes a text-autokey mechanism.
	TextAutokey

	// KeyAutokey denotes a key-autokey mechanism.
	KeyAutokey
)

// TabulaRectaCipher implements a Vigenère cipher.
type TabulaRectaCipher struct {
	*Config

	key         string
	tabulaRecta map[rune]*masc.SimpleCipher

	autokeyer autokeyFunc
}

func NewTabulaRectaCipher(key string, ciphers []*masc.SimpleCipher, autokeyer autokeyFunc, opts ...ConfigOption) (*TabulaRectaCipher, error) {
	c := NewConfig(opts...)

	tabulaRecta, err := sliceutil.Zipmap([]rune(c.keyAlphabet), ciphers)
	if err != nil {
		return nil, fmt.Errorf("could not zip alphabet with ciphers: %w", err)
	}

	return &TabulaRectaCipher{
		c,
		key,
		tabulaRecta,
		autokeyer,
	}, nil
}

// Encipher a string.
func (c *TabulaRectaCipher) Encipher(s string) (string, error) {
	if utf8.RuneCountInString(c.key) == 0 {
		return s, nil
	}

	tableau := c.tabulaRecta

	keyRunes := []rune(c.key)

	var transcodedCharCount = 0
	return strings.Map(func(r rune) rune {
		k := keyRunes[transcodedCharCount%len(keyRunes)]
		m, ok := tableau[k]
		if !ok && c.caseless {
			m, ok = tableau[unicode.ToUpper(k)]
		}
		if !ok && c.caseless {
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
			if c.autokeyer != nil {
				keyRunes = append(keyRunes, c.autokeyer(k, o, r))
			}
		}
		return o
	}, s), nil
}

// Decipher a string.
func (c *TabulaRectaCipher) Decipher(s string) (string, error) {
	if utf8.RuneCountInString(c.key) == 0 {
		return s, nil
	}

	tableau := c.tabulaRecta

	keyRunes := []rune(c.key)

	var transcodedCharCount = 0
	return strings.Map(func(r rune) rune {
		k := keyRunes[transcodedCharCount%len(keyRunes)]
		m, ok := tableau[k]
		if !ok && c.caseless {
			m, ok = tableau[unicode.ToUpper(k)]
		}
		if !ok && c.caseless {
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
			if c.autokeyer != nil {
				keyRunes = append(keyRunes, c.autokeyer(k, r, o))
			}
		}
		return o
	}, s), nil
}

// Printable representation of the tabula recta for this cipher.
func (c *TabulaRectaCipher) Printable() (string, error) {
	ptAlphabet := c.ptAlphabet

	keyAlphabet := c.keyAlphabet
	if keyAlphabet == "" {
		keyAlphabet = ptAlphabet
	}

	rt := c.tabulaRecta

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
