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

	"github.com/merenbach/goldbug/pkg/masc"
)

// NewBeaufortCipher creates and returns a new variant Beaufort cipher.
func NewVariantBeaufortCipher(key string, opts ...ConfigOption) (*TabulaRectaCipher, error) {
	c := NewConfig(opts...)

	params := []masc.ConfigOption{
		masc.WithAlphabet(c.ptAlphabet),
	}
	if c.caseless {
		params = append(params, masc.WithCaseless())
	}
	if c.strict {
		params = append(params, masc.WithStrict())
	}

	ciphers := make([]*masc.SimpleCipher, len(c.keyAlphabet))
	for i := range c.keyAlphabet {
		cipher, err := masc.NewCaesarCipher(-i, params...)
		if err != nil {
			return nil, fmt.Errorf("could not create cipher: %w", err)
		}
		ciphers[i] = cipher
	}

	return NewTabulaRectaCipher(key, ciphers, NoAutokey, opts...)
}
