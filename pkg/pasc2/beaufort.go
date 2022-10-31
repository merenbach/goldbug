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

package pasc2

import (
	"fmt"

	"github.com/merenbach/goldbug/internal/pasc"
	"github.com/merenbach/goldbug/pkg/masc"
)

// Cipher implements a Beaufort cipher.
// Cipher is effectively a Vigenère cipher with the ciphertext and key alphabets both mirrored (back-to-front).
type Cipher struct {
	*pasc.TabulaRecta
}

// adapted from: https://www.sohamkamani.com/golang/options-pattern/

func NewBeaufortCipher(key string, opts ...ConfigOption) (*Cipher, error) {
	c := NewConfig(opts...)

	// ctAlphabet, err := Transform([]rune(c.alphabet))
	// if err != nil {
	// 	return nil, fmt.Errorf("could not transform alphabet: %w", err)
	// }

	tableau, err := pasc.NewTabulaRecta(
		pasc.WithCaseless(c.caseless),
		pasc.WithPtAlphabet(c.alphabet),
		pasc.WithKeyAlphabet(c.alphabet),
		pasc.WithKey(key),
		// pasc.WithCtAlphabet(string(ctAlphabet)),
		// pasc.WithStrict(c.strict),
		pasc.WithDictFunc(func(s string, i int) (*masc.SimpleCipher, error) {
			params := []masc.ConfigOption{
				masc.WithAlphabet(s),
			}
			if c.caseless {
				params = append(params, masc.WithCaseless())
			}
			if c.strict {
				params = append(params, masc.WithStrict())
			}
			return masc.NewAffineCipher(-1, i, params...)
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("could not create tableau: %w", err)
	}

	return &Cipher{
		tableau,
	}, nil
}
