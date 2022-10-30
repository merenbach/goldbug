// Copyright 2020 Andrew Merenbach
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

package gromark

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/merenbach/goldbug/internal/lfg"
	"github.com/merenbach/goldbug/internal/masc"
	"github.com/merenbach/goldbug/internal/pasc"
	"github.com/merenbach/goldbug/pkg/caesar"
	"github.com/merenbach/goldbug/pkg/keyword"
	"github.com/merenbach/goldbug/pkg/simple"
	"github.com/merenbach/goldbug/pkg/transposition"
)

func chainadder(m int, count int, primer []int) ([]int, error) {
	g := lfg.Additive{
		Seed:    primer,
		Taps:    []int{1, 2},
		Modulus: 10,
	}

	out := make([]int, len(primer))
	copy(out, primer)
	sliced, err := g.Slice(count)
	if err != nil {
		return nil, err
	}
	out = append(out, sliced...)
	return out, nil
}

// msglen does not need to be exact; it simply needs to meet the minimum key character needs of the cipher
func makekey(k string, msglen int) (string, error) {
	primer := make([]int, utf8.RuneCountInString(k))
	for i, r := range k {
		primer[i] = int(r - '0')
	}

	raw, err := chainadder(10, msglen, primer)
	if err != nil {
		return "", err
	}

	var b strings.Builder
	for _, r := range raw {
		b.WriteRune(rune(r + '0'))
	}

	return b.String(), nil
}

// Cipher implements a GROMARK (GROnsfeld with Mixed Alphabet and Running Key) cipher.
// Cipher key is simply the primer for a running key.
type Cipher struct {
	alphabet string
	caseless bool
	key      string
	strict   bool

	*pasc.TabulaRecta
}

// adapted from: https://www.sohamkamani.com/golang/options-pattern/

type CipherOption func(*Cipher)

func WithStrict() CipherOption {
	return func(c *Cipher) {
		c.strict = true
	}
}

func WithCaseless() CipherOption {
	return func(c *Cipher) {
		c.caseless = true
	}
}

func WithAlphabet(s string) CipherOption {
	return func(c *Cipher) {
		c.alphabet = s
	}
}

func WithKey(s string) CipherOption {
	return func(c *Cipher) {
		c.key = s
	}
}

func NewCipher(key string, primer string, opts ...CipherOption) (*Cipher, error) {
	const digits = "0123456789"

	c := &Cipher{alphabet: masc.Alphabet}
	for _, opt := range opts {
		opt(c)
	}

	ctAlphabetInput, _ := keyword.Transform(c.alphabet, key)

	tc := transposition.Cipher{
		Keys: []string{key},
	}
	transposedCtAlphabet, err := tc.Encipher(ctAlphabetInput)
	if err != nil {
		return nil, err
	}

	// t, err := masc.NewTableau(ptAlphabet, transposedCtAlphabet, func(string) (string, error) {
	// 	return transposedCtAlphabet, nil
	// })
	params := []pasc.TabulaRectaOption{
		pasc.WithPtAlphabet(c.alphabet),
		pasc.WithKeyAlphabet(digits),
		pasc.WithKey(key),
		pasc.WithKeyGenerator(func(s string) (string, error) {
			return makekey(primer, utf8.RuneCountInString(s))
		}),
		pasc.WithDictFunc(func(s string, i int) (*masc.Tableau, error) {
			ctAlphabetTransformed, err := caesar.Transform([]rune(transposedCtAlphabet), i)
			if err != nil {
				return nil, fmt.Errorf("could not transform alphabet: %w", err)
			}

			params := []simple.CipherOption{
				simple.WithPtAlphabet(s),
				simple.WithCtAlphabet(string(ctAlphabetTransformed)),
			}
			if c.caseless {
				params = append(params, simple.WithCaseless())
			}
			if c.strict {
				params = append(params, simple.WithStrict())
			}
			c2, err := simple.NewCipher(params...)
			if err != nil {
				return nil, fmt.Errorf("could not create cipher: %w", err)
			}
			return c2.Tableau, nil
		}),
	}
	if c.caseless {
		params = append(params, pasc.WithCaseless(true))
	}

	tableau, err := pasc.NewTabulaRecta(params...)
	if err != nil {
		return nil, fmt.Errorf("could not create tableau: %w", err)
	}
	c.TabulaRecta = tableau

	return c, nil
}

type Cipher2 struct {
	Alphabet string
	Caseless bool
	Key      string
	Primer   string
	Strict   bool
}

// // ALLOW COMPOUND CIPHER CHAINING:
// // NewCompoundCipher(keyword.Cipher, transposition.Cipher, caesar.Cipher...).Tableau()
// type cipher interface {
// 	Encipher(string) (string, error)
// 	Decipher(string) (string, error)
// 	// Tableau(string) (*masc.Tableau, error)
// }

// func chainedTableauEncipherment(s string, cc ...cipher) (string, error) {
// 	var err error
// 	for _, c := range cc {
// 		s, err = c.Encipher(s)
// 		if err != nil {
// 			return "", err
// 		}
// 	}
// 	return s, nil
// }

// // Encipher a message.
// func (c *Cipher) Encipher(s string) (string, error) {
// 	t := c.TabulaRecta

// 	key, err := makekey(c.primer, utf8.RuneCountInString(s))
// 	if err != nil {
// 		return "", err
// 	}
// 	return t.Encipher(s, key, nil)
// }

// // Decipher a message.
// func (c *Cipher) Decipher(s string) (string, error) {
// 	t := c.TabulaRecta

// 	key, err := makekey(c.primer, utf8.RuneCountInString(s))
// 	if err != nil {
// 		return "", err
// 	}
// 	return t.Decipher(s, key, nil)
// }

// // Tableau for encipherment and decipherment.
// func (c *Cipher) Tableau() (string, error) {
// 	t := c.TabulaRecta
// 	return t.Printable()
// }
