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

package vigenere

import (
	"github.com/merenbach/goldbug/internal/pasc"
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

func autokeyOptionFromString(s string) autokeyOption {
	switch s {
	case "TextAutokey":
		return TextAutokey
	case "KeyAutokey":
		return KeyAutokey
	default:
		return NoAutokey
	}
}

// Cipher implements a Vigenere cipher.
type Cipher struct {
	Alphabet string
	Autokey  autokeyOption
	Key      string
	Strict   bool
}

// NewCipher creates a new Vigenere cipher.
func newCipher(countersign string, alphabet string) (*pasc.VigenereFamilyCipher, error) {
	if alphabet == "" {
		alphabet = pasc.Alphabet
	}
	return pasc.NewVigenereFamilyCipher(countersign, alphabet, alphabet, alphabet)
}

// Encipher a message.
func (c *Cipher) Encipher(s string) (string, error) {
	c2, err := newCipher(c.Key, c.Alphabet)
	if err != nil {
		return "", err
	}

	c2.Strict = c.Strict
	switch c.Autokey {
	case TextAutokey:
		c2.TextAutoclave = true
	case KeyAutokey:
		c2.KeyAutoclave = true
	}
	return c2.EncipherString(s), nil
}

// Decipher a message.
func (c *Cipher) Decipher(s string) (string, error) {
	c2, err := newCipher(c.Key, c.Alphabet)
	if err != nil {
		return "", err
	}

	c2.Strict = c.Strict
	switch c.Autokey {
	case TextAutokey:
		c2.TextAutoclave = true
	case KeyAutokey:
		c2.KeyAutoclave = true
	}
	return c2.DecipherString(s), nil
}

// Tableau for encipherment and decipherment.
func (c *Cipher) tableau() (string, error) {
	c2, err := newCipher(c.Key, c.Alphabet)
	if err != nil {
		return "", err
	}

	c2.Strict = c.Strict
	return c2.Printable(), nil
}
