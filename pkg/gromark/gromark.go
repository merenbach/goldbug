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
	"unicode/utf8"

	"github.com/merenbach/goldbug/internal/lfg"
	"github.com/merenbach/goldbug/internal/pasc"
	"github.com/merenbach/goldbug/internal/sliceutil"
	"github.com/merenbach/goldbug/pkg/masc"
	"github.com/merenbach/goldbug/pkg/transposition"
)

// msglen does not need to be exact; it simply needs to meet the minimum key character needs of the cipher
func makekey(k string) (func() rune, error) {
	primer := make([]int, utf8.RuneCountInString(k))
	for i, r := range k {
		primer[i] = int(r - '0')
	}

	g := lfg.Additive{
		Seed:    primer,
		Taps:    []int{1, 2},
		Modulus: 10,
	}

	iter, err := g.Iterate()
	if err != nil {
		return nil, fmt.Errorf("couldn't configure chain adder: %w", err)
	}

	return func() rune {
		return rune(iter() + '0')
	}, nil
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

	c := &Cipher{alphabet: pasc.Alphabet}
	for _, opt := range opts {
		opt(c)
	}

	ctAlphabetInput, _ := sliceutil.Keyword([]rune(c.alphabet), []rune(key))

	tc := transposition.Cipher{
		Keys: []string{key},
	}
	transposedCtAlphabet, err := tc.Encipher(string(ctAlphabetInput))
	if err != nil {
		return nil, err
	}

	// t, err := masc.NewTableau(ptAlphabet, transposedCtAlphabet, func(string) (string, error) {
	// 	return transposedCtAlphabet, nil
	// })
	keygen, err := makekey(primer)
	if err != nil {
		return nil, fmt.Errorf("couldn't create running key generator: %w", err)
	}

	params := []pasc.TabulaRectaOption{
		pasc.WithPtAlphabet(c.alphabet),
		pasc.WithKeyAlphabet(digits),
		pasc.WithKey(primer),
		pasc.WithAutokeyer(func(_ rune, _ rune, keystream *[]rune) {
			*keystream = append(*keystream, keygen())
		}),
		pasc.WithDictFunc(func(s string, i int) (*masc.SimpleCipher, error) {
			ctAlphabetTransformed, err := sliceutil.Affine([]rune(transposedCtAlphabet), 1, i)
			if err != nil {
				return nil, fmt.Errorf("could not transform alphabet: %w", err)
			}

			params := []masc.ConfigOption{
				masc.WithAlphabet(s),
			}
			if c.caseless {
				params = append(params, masc.WithCaseless())
			}
			if c.strict {
				params = append(params, masc.WithStrict())
			}
			return masc.NewSimpleCipher(string(ctAlphabetTransformed), params...)
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
