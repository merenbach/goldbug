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

package atbash

import "github.com/merenbach/goldbug/pkg/affine"

// Intercept for the affine cipher.
const intercept = (-1)

// Slope for the affine cipher.
const slope = (-1)

// Cipher implements an Atbash cipher.
type Cipher struct {
	Alphabet string
	Strict   bool
}

// Encipher a message.
func (c *Cipher) Encipher(s string) (string, error) {
	c2 := affine.Cipher{
		Alphabet:  c.Alphabet,
		Intercept: intercept,
		Slope:     slope,
		Strict:    c.Strict,
	}
	return c2.Encipher(s)
}

// Decipher a message.
func (c *Cipher) Decipher(s string) (string, error) {
	c2 := affine.Cipher{
		Alphabet:  c.Alphabet,
		Intercept: intercept,
		Slope:     slope,
		Strict:    c.Strict,
	}
	return c2.Decipher(s)
}

// Tableau for encipherment and decipherment.
func (c *Cipher) Tableau() (string, error) {
	c2 := affine.Cipher{
		Alphabet:  c.Alphabet,
		Intercept: intercept,
		Slope:     slope,
		Strict:    c.Strict,
	}
	return c2.Tableau()
}
