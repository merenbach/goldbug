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

package affine

import (
	"fmt"

	"github.com/merenbach/goldbug/internal/masc"
)

// Cipher implements an affine cipher.
type Cipher struct {
	Alphabet   string
	CtAlphabet string
	Caseless   bool
	Intercept  int
	Slope      int
	Strict     bool
}

func (c *Cipher) maketableau() (*masc.Tableau, error) {
	config := &masc.Configuration{
		Alphabet: c.Alphabet,
		Strict:   c.Strict,
		Caseless: c.Caseless,
	}
	t, err := masc.NewTableau(config, c.CtAlphabet, func(s string) (string, error) {
		out, err := transform([]rune(s), c.Slope, c.Intercept)
		if err != nil {
			return "", err
		}
		return string(out), nil
	})
	if err != nil {
		return nil, err
	}
	return t, nil
}

// Encipher a message.
func (c *Cipher) Encipher(s string) (string, error) {
	t, err := c.maketableau()
	if err != nil {
		return "", fmt.Errorf("could not calculate encipherment alphabets: %w", err)
	}
	return t.Encipher(s)
}

// Decipher a message.
func (c *Cipher) Decipher(s string) (string, error) {
	t, err := c.maketableau()
	if err != nil {
		return "", fmt.Errorf("could not calculate decipherment alphabets: %w", err)
	}
	return t.Decipher(s)
}

// Tableau for this cipher.
func (c *Cipher) Tableau() (*masc.Tableau, error) {
	return c.maketableau()
}
