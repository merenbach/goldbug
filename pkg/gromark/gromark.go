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
	"strings"
	"unicode/utf8"

	"github.com/merenbach/goldbug/internal/lfg"
	"github.com/merenbach/goldbug/internal/masc"
	"github.com/merenbach/goldbug/internal/pasc"
	"github.com/merenbach/goldbug/internal/stringutil"
	"github.com/merenbach/goldbug/pkg/caesar"
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
	Alphabet string
	// Caseless bool
	Key    string
	Primer string
	Strict bool
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

func (c *Cipher) maketableau() (*pasc.TabulaRecta, error) {
	const digits = "0123456789"

	ptAlphabet := c.Alphabet
	if ptAlphabet == "" {
		ptAlphabet = pasc.Alphabet
	}

	// 	c1 := keyword.Cipher{
	// 		Alphabet: ptAlphabet,
	// 		// Caseless: ...,
	// 		Keyword: c.Key,
	// 		// Strict: ...,
	// 	}
	// 	c2 := transposition.Cipher{
	// 		Keys: []string{c.Key},
	// 	}
	// 	c3 := caesar.Cipher
	// ...
	ctAlphabetInput := stringutil.Key(ptAlphabet, c.Key)

	tc := transposition.Cipher{
		Keys: []string{c.Key},
	}
	transposedCtAlphabet, err := tc.Encipher(ctAlphabetInput)
	if err != nil {
		return nil, err
	}

	// t, err := masc.NewTableau(ptAlphabet, transposedCtAlphabet, func(string) (string, error) {
	// 	return transposedCtAlphabet, nil
	// })
	tr, err := pasc.NewTabulaRecta(c.Alphabet, digits, func(s string, i int) (*masc.Tableau, error) {
		c2 := &caesar.Cipher{
			Alphabet:   s,
			CtAlphabet: transposedCtAlphabet,
			Shift:      i,
		}
		return c2.Tableau()
	})
	if err != nil {
		return nil, err
	}

	// tr.Caseless = c.Caseless
	tr.Strict = c.Strict
	return tr, nil
}

// Encipher a message.
func (c *Cipher) Encipher(s string) (string, error) {
	t, err := c.maketableau()
	if err != nil {
		return "", err
	}

	key, err := makekey(c.Primer, utf8.RuneCountInString(s))
	if err != nil {
		return "", err
	}
	return t.Encipher(s, key, nil)
}

// Decipher a message.
func (c *Cipher) Decipher(s string) (string, error) {
	t, err := c.maketableau()
	if err != nil {
		return "", err
	}

	key, err := makekey(c.Primer, utf8.RuneCountInString(s))
	if err != nil {
		return "", err
	}
	return t.Decipher(s, key, nil)
}

// Tableau for encipherment and decipherment.
func (c *Cipher) Tableau() (string, error) {
	t, err := c.maketableau()
	if err != nil {
		return "", err
	}
	return t.Printable()
}
