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
	*pasc.TabulaRecta
}

// adapted from: https://www.sohamkamani.com/golang/options-pattern/

func NewTabulaRectaCipher(key string, ciphers []*masc.SimpleCipher, autokey autokeyOption, opts ...ConfigOption) (*TabulaRectaCipher, error) {
	c := NewConfig(opts...)

	params := []pasc.TabulaRectaOption{
		pasc.WithCaseless(c.caseless),
		pasc.WithPtAlphabet(c.ptAlphabet),
		pasc.WithKeyAlphabet(c.keyAlphabet),
		pasc.WithKey(key),
		// pasc.WithCtAlphabet(string(ctAlphabet)),
		// pasc.WithStrict(c.strict),
		pasc.WithDictFunc(func(s string, i int) (*masc.SimpleCipher, error) {
			return ciphers[i%len(ciphers)], nil
		}),
	}
	params = append(params, pasc.WithAutokeyer(func(a rune, b rune, keystream *[]rune) {
		switch autokey {
		case NoAutokey:
			// do nothing
		case KeyAutokey:
			*keystream = append(*keystream, b)
		case TextAutokey:
			*keystream = append(*keystream, a)
		}
	}))

	tableau, err := pasc.NewTabulaRecta(params...)
	if err != nil {
		return nil, fmt.Errorf("could not create tableau: %w", err)
	}

	return &TabulaRectaCipher{tableau}, nil
}
