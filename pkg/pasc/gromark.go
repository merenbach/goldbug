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

package pasc

import (
	"fmt"
	"unicode/utf8"

	"github.com/merenbach/goldbug/internal/lfg"
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

// NewGromarkCipher creates and returns a new GROMARK (GROnsfeld with Mixed Alphabet and Running Key) cipher.
func NewGromarkCipher(key string, primer string, opts ...ConfigOption) (*TabulaRectaCipher, error) {
	const digits = "0123456789"

	opts = append(opts, WithKeyAlphabet(digits))
	c := NewConfig(opts...)

	ctAlphabetInput, _ := sliceutil.Keyword([]rune(c.ptAlphabet), []rune(key))

	tc := transposition.NewCipher([]string{key})
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

	ciphers := make([]*masc.SimpleCipher, utf8.RuneCountInString(digits))
	for i := range ciphers {
		ctAlphabetTransformed, err := sliceutil.Affine([]rune(transposedCtAlphabet), 1, i)
		if err != nil {
			return nil, fmt.Errorf("could not transform alphabet: %w", err)
		}

		s := c.ptAlphabet

		params := []masc.ConfigOption{
			masc.WithAlphabet(s),
		}
		if c.caseless {
			params = append(params, masc.WithCaseless())
		}
		if c.strict {
			params = append(params, masc.WithStrict())
		}
		cipher, err := masc.NewSimpleCipher(string(ctAlphabetTransformed), params...)
		if err != nil {
			return nil, fmt.Errorf("could not create cipher: %w", err)
		}
		ciphers[i] = cipher
	}

	return NewTabulaRectaCipher(primer, ciphers, func(_ rune, _ rune, _ rune) rune {
		return keygen()
	}, opts...)
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