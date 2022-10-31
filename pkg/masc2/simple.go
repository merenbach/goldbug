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

package masc2

import (
	"fmt"
)

// A Cipher implements a simple cipher.
type Cipher struct {
	*Tableau
}

// NewSimpleCipher creates and returns a new simple cipher.
func NewSimpleCipher(ctAlphabet string, opts ...ConfigOption) (*Cipher, error) {
	c := NewConfig(opts...)

	tableau, err := NewTableau(
		c.Alphabet(),
		ctAlphabet,
		c.Strict(),
		c.Caseless(),
	)
	if err != nil {
		return nil, fmt.Errorf("could not create tableau: %w", err)
	}

	return &Cipher{tableau}, nil
}
