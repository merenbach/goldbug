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
	"unicode/utf8"

	"github.com/merenbach/goldbug/internal/masc"
	"github.com/merenbach/goldbug/internal/stringutil"
)

// TabulaRecta holds a tabula recta.
type TabulaRecta struct {
	Strict bool

	PtAlphabet  string
	CtAlphabet  string
	KeyAlphabet string

	DictFunc func(s string, i int) (*masc.Tableau, error)
}

// MakeTabulaRecta creates a standard Caesar shift tabula recta.
func (tr *TabulaRecta) makereciprocaltable() (*ReciprocalTable, error) {
	ctAlphabets := make([]string, utf8.RuneCountInString(tr.KeyAlphabet))
	ctAlphabetLen := utf8.RuneCountInString(tr.CtAlphabet)

	// Cast to []rune to increase index without gaps
	for y := range ctAlphabets {
		ii := make([]int, ctAlphabetLen)
		for x := range ii {
			ii[x] = (x + y) % ctAlphabetLen
		}

		out, err := stringutil.Backpermute(tr.CtAlphabet, ii)
		if err != nil {
			return nil, err
		}
		ctAlphabets[y] = out
	}

	rt := ReciprocalTable{
		PtAlphabet:  tr.PtAlphabet,
		KeyAlphabet: tr.KeyAlphabet,
		DictFunc:    tr.DictFunc,
		CtAlphabets: ctAlphabets,
		Strict:      tr.Strict,
	}

	return &rt, nil
}

func (tr *TabulaRecta) String() string {
	return fmt.Sprintf("%+v", map[string]interface{}{
		"k":  tr.KeyAlphabet,
		"pt": tr.PtAlphabet,
		"ct": tr.CtAlphabet,
	})
}

// Printable representation of this tabula recta.
func (tr *TabulaRecta) Printable() (string, error) {
	rt, err := tr.makereciprocaltable()
	if err != nil {
		return "", err
	}
	return rt.Printable()
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
func (tr *TabulaRecta) Encipher(s string, k string, onSuccess func(rune, rune, *[]rune)) (string, error) {
	rt, err := tr.makereciprocaltable()
	if err != nil {
		return "", err
	}
	return rt.Encipher(s, k, onSuccess)
}

// Decipher a string.
func (tr *TabulaRecta) Decipher(s string, k string, onSuccess func(rune, rune, *[]rune)) (string, error) {
	rt, err := tr.makereciprocaltable()
	if err != nil {
		return "", err
	}
	return rt.Decipher(s, k, onSuccess)
}
