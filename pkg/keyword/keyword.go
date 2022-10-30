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

package keyword

import (
	"fmt"

	"github.com/merenbach/goldbug/internal/masc"
	"github.com/merenbach/goldbug/pkg/affine"
)

// A Cipher implements a keyword cipher.
type Cipher struct {
	*affine.Config
	*masc.Tableau
}

func NewCipher(keyword string, opts ...affine.ConfigOption) (*Cipher, error) {
	c := affine.NewConfig(opts...)
	ctAlphabet, _ := Transform([]rune(c.Alphabet()), []rune(keyword))

	tableau, err := masc.NewTableau(
		c.Alphabet(),
		string(ctAlphabet),
		c.Strict(),
		c.Caseless(),
	)
	if err != nil {
		return nil, fmt.Errorf("could not create tableau: %w", err)
	}

	return &Cipher{c, tableau}, nil
}
