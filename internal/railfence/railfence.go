// Copyright 2019 Andrew Merenbach
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

package railfence

import (
	"unicode/utf8"
)

// A Cipher implements the rail fence (or zig-zag) cipher.
type Cipher struct {
	Rails int
}

// Encipher a message.
func (c *Cipher) encipher(s string) *grid {
	cc := make(grid, utf8.RuneCountInString(s))
	cc.renumber(c.Rails)

	cc.fill(s)
	cc.sortByRow()

	return &cc
}

// Decipher a message.
func (c *Cipher) decipher(s string) *grid {
	cc := make(grid, utf8.RuneCountInString(s))
	cc.renumber(c.Rails)

	cc.sortByRow()
	cc.fill(s)
	cc.sortByCol()

	return &cc
}

// Encipher a message.
func (c *Cipher) Encipher(s string) string {
	if c.Rails == 1 {
		return s
	}
	return c.encipher(s).contents()
}

// Decipher a message.
func (c *Cipher) Decipher(s string) string {
	if c.Rails == 1 {
		return s
	}
	return c.decipher(s).contents()
}

// EnciphermentGrid returns the output tableau upon encipherment.
func (c *Cipher) EnciphermentGrid(s string) string {
	if c.Rails == 1 {
		return s
	}
	return c.encipher(s).printable()
}

// DeciphermentGrid returns the output tableau upon encipherment.
func (c *Cipher) DeciphermentGrid(s string) string {
	if c.Rails == 1 {
		return s
	}
	return c.decipher(s).printable()
}
